package pdftools

import (
	"go_recipes/models"
	"go_recipes/repository"
	"go_recipes/utils"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/textract"
	"github.com/joho/godotenv"
)

type pdfReader struct {
	paths             PdfToImport
	ingredientsPages  []uint16
	ingredientstoSkip []string
	stepsPages        []uint16
	stepsToSkip       []string
}

// textract (pdf reader)
var textractSession *textract.Textract

// logger
var logger utils.Logger = utils.NewLogger("reader.log")

// db connection
var sql *repository.MySQLRepository = repository.DbConnect(utils.GetEnvFile().Name)

func init() {
	// start Textract session
	envFile := utils.GetEnvFile().Name
	if err := godotenv.Load(envFile); err != nil {
		panic(err)
	}

	textractSession = textract.New(session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	})))
}

func newReader(ingPages []uint16, stepPages []uint16, ingtoSkip []string, pdfPaths PdfToImport) pdfReader {
	return pdfReader{
		paths:             pdfPaths,
		ingredientsPages:  ingPages,
		stepsPages:        stepPages,
		ingredientstoSkip: ingtoSkip,
	}
}

func (reader pdfReader) read() error {

	dish, err := reader.readDish()
	if err != nil {
		return err
	}
	reader.readQuantities(dish)
	// TODO read steps
	return nil

}

func (reader pdfReader) readDish() (*models.Dish, error) {
	dishRepo := repository.DishRepository{Mysql: *sql}
	// User filename as Dish name
	dish, err := dishRepo.GetOrCreate(reader.paths.FileName)
	if err != nil {
		return nil, err
	}
	return dish, nil
}

func (reader pdfReader) readQuantities(dish *models.Dish) {
	// read all pages with ingredients
	var wg sync.WaitGroup
	for _, page := range reader.ingredientsPages {

		fileName := reader.paths.Path + reader.paths.FileName + "-" +
			strconv.FormatInt(int64(page), 10) + reader.paths.Extension

		wg.Add(1)
		go reader.getQuantities(fileName, dish)
	}
	wg.Wait()
	wg.Done()
}

func (reader pdfReader) getQuantities(fileName string, dish *models.Dish) error {

	resp, err := openFile(fileName)
	if err != nil {
		logger.Sugar.Fatal(err)
	}

	// count read lines
	lines := 1

	for i := 1; i < len(resp.Blocks); i++ {
		if i <= 7 {
			// skip the useless informations from pdf
			// too specific, must be a parameter
			continue
		}

		if lines%2 == 0 {
			// skip lines with quantities to create directly ingredients + quantity with i+1
			lines++
			continue
		}

		if *resp.Blocks[i].BlockType == "LINE" {
			if isUselessData(resp.Blocks[i].Text, reader.ingredientstoSkip) {
				continue
			}

			// get or create ing
			ingRepo := repository.IngRepository{Mysql: *sql}
			ingDb, err := ingRepo.GetOrCreate(*resp.Blocks[i].Text)
			if err != nil {
				return err
			}

			qtt, err := getQuantityEntity(resp.Blocks[i+1].Text, dish.ID, ingDb.ID)
			if err != nil {
				return err
			}

			// get or create qtt
			qttRepo := repository.QuantityRepository{Mysql: *sql}
			_, qttErr := qttRepo.GetOrCreate(qtt)
			if qttErr != nil {
				return qttErr
			}

			// push qtt (with ing) in dish
			dish.Quantities = append(dish.Quantities, qtt)
			/* fmt.Println(*resp.Blocks[i].Text) */
			lines++
		}

	}

	return nil
}

func getQuantityEntity(data *string, dishId uint, ingId uint) (models.Quantity, error) {
	var entity models.Quantity
	qtt := strings.Fields(*data)
	qttInt, err := strconv.Atoi(qtt[0])
	if err != nil {
		return entity, err
	}

	entity = models.Quantity{
		Num:          uint16(qttInt),
		Unit:         qtt[1],
		DishID:       dishId,
		IngredientID: ingId,
	}

	return entity, nil
}

func openFile(fileName string) (*textract.DetectDocumentTextOutput, error) {
	// Open file with AWS textract to read it
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	resp, err := textractSession.DetectDocumentText(&textract.DetectDocumentTextInput{
		Document: &textract.Document{
			Bytes: file,
		},
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func isUselessData(input *string, uselesses []string) bool {
	// skip useless data from pdf
	for _, v := range uselesses {
		if v == *input {
			return true
		}
	}
	return false
}
