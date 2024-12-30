package main

import (
	pdftools "go_recipes/utils/pdf_tools"
	"log"
)

func main() {

	pdf := pdftools.PdfToImport{
		FileName:  "test1",
		Extension: ".pdf",
		Path:      "./assets/",
	}

	err := pdf.CropFile(0, 0, -50, 0)
	if err != nil {
		log.Fatalln(err)
	}

	splitErr := pdf.SplitFile(4, 2)
	if splitErr != nil {
		panic(splitErr)
	}

	/* 	cmd := exec.Command("echo", "hello", "world")
	   	if errors.Is(cmd.Err, exec.ErrDot) {
	   		cmd.Err = nil
	   	} */

	/* 	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	} */
	/* 	stdout, outErr := cmd.Output()
	   	if outErr != nil {
	   		log.Fatalln(outErr)
	   	}

	   	fmt.Println(string(stdout))
	*/
	/*
		 	err := repository.Migrate()
			if err != nil {
				log.Fatalln(err)
			}
	*/
}
