package pdftools

import (
	"fmt"
	"go_recipes/models"
	"go_recipes/utils"
	"os"
	"strconv"
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

var textractSession *textract.Textract
var logger utils.Logger = utils.NewLogger("reader.log")

func init() {
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

	reader.readQuantities()
	return nil

}

func (reader pdfReader) readQuantities() {
	// read all pages with ingredients
	var quantities []models.Quantity
	var wg sync.WaitGroup
	for _, page := range reader.ingredientsPages {

		fileName := reader.paths.Path + reader.paths.FileName + "-" +
			strconv.FormatInt(int64(page), 10) + reader.paths.Extension

		wg.Add(1)
		go reader.getQuantities(fileName, &quantities)
	}
	wg.Wait()
}

func (reader pdfReader) getQuantities(fileName string, quantities *[]models.Quantity) error {
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
		if *resp.Blocks[i].BlockType == "LINE" {
			if isUselessData(resp.Blocks[i].Text, reader.ingredientstoSkip) {
				continue
			}

			// TODO sort ingredients and quantities and create entities

			fmt.Println(*resp.Blocks[i].Text)
			lines++
		}
	}

	return nil
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
