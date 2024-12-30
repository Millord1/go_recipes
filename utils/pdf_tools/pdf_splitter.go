package pdftools

import (
	"errors"
	"os/exec"
)

type pdfSplitter struct {
	cmd         string
	arg         string
	XDecimation uint16
	YDecimation uint16
	paths       PdfToImport
}

func newSplitter(XDecimation uint16, YDecimation uint16, pdfPaths PdfToImport) pdfSplitter {
	return pdfSplitter{
		cmd:         "mutool",
		arg:         "poster",
		XDecimation: XDecimation,
		YDecimation: YDecimation,
		paths:       pdfPaths,
	}
}

func (splitter pdfSplitter) split() error {
	splitErr := splitter.splitPdf()
	if splitErr != nil {
		return splitErr
	}

	sepErr := splitter.separatePages()
	if sepErr != nil {
		return sepErr
	}
	return nil
}

func (splitter pdfSplitter) splitPdf() error {
	// split pdf by height and width, depending of XDecimation and YDecimation
	// generate pages into single file pdf
	fileName := splitter.paths.getFullFileName()
	cmd := exec.Command(splitter.cmd, splitter.arg,
		"-x "+intToString(int16(splitter.XDecimation)),
		"-y "+intToString(int16(splitter.YDecimation)),
		fileName, fileName)

	stdOut, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(stdOut))
	}
	return nil
}

func (splitter pdfSplitter) separatePages() error {
	// separate single file pdf pages into single page pdf files
	separate := exec.Command("pdfseparate", splitter.paths.getFullFileName(),
		splitter.paths.getPathAndName()+"-%d"+splitter.paths.Extension)

	splitOut, splitErr := separate.CombinedOutput()
	if splitErr != nil {
		return errors.New(string(splitOut))
	}
	// TODO remove original file
	return nil
}
