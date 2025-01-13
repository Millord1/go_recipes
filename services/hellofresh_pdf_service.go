package services

import (
	pdftools "go_recipes/utils/pdf_tools"
	"path/filepath"
	"strings"
)

type HelloFreshPdf struct {
	pdf         pdftools.PdfToImport
	uselessIng  []string
	ingPages    []uint16
	uselessStep []string
	stepsPages  []uint16
	cropData    [4]int16
	splitData   [2]uint16
}

func NewHelloFreshPdf(pathToFiles string) HelloFreshPdf {

	return HelloFreshPdf{
		pdf: pdftools.PdfToImport{
			Extension: filepath.Ext(pathToFiles),
			Path:      filepath.Dir(pathToFiles) + "/",
			FileName:  strings.Split(filepath.Base(pathToFiles), ".")[0],
		},
		uselessIng: getUselessIng(),
		// pages tp read as 'ingredient'
		ingPages:    []uint16{1},
		uselessStep: getUselessSteps(),
		// pages tp read as 'step'
		stepsPages: []uint16{2, 3, 4, 6, 7, 8},
		// in order: left, top, right and bottom to crop
		cropData: [4]int16{0, 0, -50, 0},
		// first element is X decimation, second Y decimation to split
		splitData: [2]uint16{4, 2},
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
