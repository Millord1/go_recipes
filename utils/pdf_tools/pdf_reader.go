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

func newReader(ingPages []uint16, stepPages []uint16, ingtoSkip []string, stepToSkip []string, pdfPaths PdfToImport) pdfReader {
	return pdfReader{
		paths:             pdfPaths,
		ingredientsPages:  ingPages,
		stepsPages:        stepPages,
		ingredientstoSkip: ingtoSkip,
		stepsToSkip:       stepToSkip,
	}
}

func (reader pdfReader) read() error {

	dish, err := reader.readDish()
	if err != nil {
		return err
	}

	reader.readQuantities(dish)

	reader.readSteps(dish)

	dishRepo := repository.DishRepository{Mysql: *sql}
	updateErr := dishRepo.Update(dish)
	if updateErr != nil {
		return updateErr
	}
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

		fileName := reader.getPageName(page)

		wg.Add(1)
		go func() {
			reader.getQuantities(fileName, dish)
			defer wg.Done()
		}()
	}
	wg.Wait()
}

func (reader pdfReader) readSteps(dish *models.Dish) {
	var wg sync.WaitGroup
	for _, page := range reader.stepsPages {
		fileName := reader.getPageName(page)
		wg.Add(1)
		go func() {
			reader.getSteps(fileName, dish)
			wg.Done()
		}()
	}
	wg.Wait()
}

func (reader pdfReader) getSteps(fileName string, dish *models.Dish) error {
	resp, err := openFile(fileName)

	if err != nil {
		logger.Sugar.Fatal(err)
		return err
	}

	var sb strings.Builder
	var order int
	var title string

	for i := 0; i < len(resp.Blocks); i++ {
		if *resp.Blocks[i].BlockType == "LINE" {

			if i == 1 {
				var convErr error
				order, convErr = strconv.Atoi(*resp.Blocks[i].Text)
				if convErr != nil {
					logger.Sugar.Error(err)
					return err
				}
				continue
			}

			if i == 2 {
				title = *resp.Blocks[i].Text
				continue
			}

			if isUselessData(*resp.Blocks[i].Text, reader.stepsToSkip) {
				continue
			}
			// All the remaining lines are just content, so we concatenate them
			sb.WriteString(*resp.Blocks[i].Text)
		}
	}

	step := models.Step{
		Order:   order,
		Title:   title,
		Content: sb.String(),
	}

	dish.Steps = append(dish.Steps, step)

	return nil
}

func (reader pdfReader) getQuantities(fileName string, dish *models.Dish) error {

	resp, err := openFile(fileName)
	if err != nil {
		logger.Sugar.Fatal(err)
		return err
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
			if isUselessData(*resp.Blocks[i].Text, reader.ingredientstoSkip) {
				continue
			}

			// get or create ing
			ingRepo := repository.IngRepository{Mysql: *sql}

			ingName := strings.Replace(*resp.Blocks[i].Text, "*", "", -1)
			ingDb, err := ingRepo.GetOrCreate(ingName)
			if err != nil {
				return err
			}

			qtt, err := getQuantityEntity(resp.Blocks[i+1].Text, dish.ID, ingDb.ID)
			if err != nil {
				return err
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

func (reader pdfReader) getPageName(page uint16) string {
	return reader.paths.Path + reader.paths.FileName + "-" +
		strconv.FormatInt(int64(page), 10) + reader.paths.Extension
}

func isUselessData(input string, uselesses []string) bool {
	// skip useless data from pdf
	for _, v := range uselesses {
		if v == input {
			return true
		}
	}
	return false
}
