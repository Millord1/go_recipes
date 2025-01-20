package services

import (
	"fmt"
	pdftools "go_recipes/utils/pdf_tools"
	"log"
	"os"
)

type HelloFreshPdf struct {
	pdf         pdftools.PdfToImport
	uselessIng  []string
	ingPages    []uint16
	uselessStep []string
	stepsPages  []uint16
	cropData    [4]int16
	splitData   [2]uint16
	path        string
}

func NewHelloFreshPdf(path string) HelloFreshPdf {

	path, err := pdftools.GetAbsPath(path)

	if err != nil {
		log.Fatalln(err)
	}

	return HelloFreshPdf{
		/* 		pdf: pdftools.PdfToImport{
			Extension: filepath.Ext(path),
			Path:      filepath.Dir(path) + "/",
			FileName:  strings.Split(filepath.Base(path), ".")[0],
		}, */
		uselessIng: getUselessIng(),
		// pages to read as 'ingredient'
		ingPages:    []uint16{1},
		uselessStep: getUselessSteps(),
		// pages to read as 'step'
		stepsPages: []uint16{2, 3, 4, 6, 7, 8},
		// in order: left, top, right and bottom to crop
		cropData: [4]int16{0, 0, -50, 0},
		// first element is X decimation, second Y decimation to split
		splitData: [2]uint16{4, 2},
		path:      path,
	}
}

func getUselessIng() []string {
	// Data that we don't want to keep
	return []string{
		"A ajouter vous-meme",
		"* Conserver au réfrigérateur",
		"Valeurs nutritionnelles",
		"Par portion Pour 100 g",
	}
}

func getUselessSteps() []string {
	// Data that we don't want to keep
	return []string{
		"LETRI",
		"+ FACILE",
		"A vos fourchettes !",
	}
}

func (hfPdf HelloFreshPdf) GetAllFilesToRead() ([]string, error) {
	files, err := os.ReadDir(hfPdf.path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Printf("Files in %s : \n", hfPdf.path)

	var allFiles []string

	for _, file := range files {
		allFiles = append(allFiles, file.Name())
		fmt.Println(file.Name(), file.IsDir())
	}

	return allFiles, nil
}
