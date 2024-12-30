package pdftools

import (
	"go_recipes/utils"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/textract"
	"github.com/joho/godotenv"
)

type pdfReader struct {
	paths PdfToImport
}

var textractSession *textract.Textract

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

func newReader(pdfPaths PdfToImport) pdfReader {
	return pdfReader{
		paths: pdfPaths,
	}
}

func (reader pdfReader) read() error {
	fileName := reader.paths.Path + reader.paths.FileName + reader.paths.Extension
	file, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	resp, err := textractSession.DetectDocumentText(&textract.DetectDocumentTextInput{
		Document: &textract.Document{
			Bytes: file,
		},
	})

	if err != nil {
		return err
	}

	log.Fatalf("%+v\n", resp)
	return nil

	// TODO skip one page and read remaining
	// create entities

}
