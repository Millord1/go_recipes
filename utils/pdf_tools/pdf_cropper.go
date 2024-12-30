package pdftools

import (
	"errors"
	"os/exec"
	"strings"
)

type pdfCropper struct {
	cmd          string
	arg          string
	MarginLeft   int16
	MarginTop    int16
	MarginRight  int16
	MarginBottom int16
	paths        PdfToImport
}

func newCropper(left int16, top int16, right int16, bottom int16, pdfPaths PdfToImport) pdfCropper {
	return pdfCropper{
		cmd:          "pdfcrop",
		arg:          "--margins",
		MarginLeft:   left,
		MarginTop:    top,
		MarginRight:  right,
		MarginBottom: bottom,
		paths:        pdfPaths,
	}
}

func (cropper pdfCropper) crop() error {
	// Crop pdf file, depending of margins
	// useful to split and read pdf

	var sb strings.Builder
	sb.WriteString(intToString(cropper.MarginLeft) + " ")
	sb.WriteString(intToString(cropper.MarginTop) + " ")
	sb.WriteString(intToString(cropper.MarginRight) + " ")
	sb.WriteString(intToString(cropper.MarginBottom))

	cmd := exec.Command(cropper.cmd, cropper.arg, sb.String(),
		cropper.paths.getFullFileName(), cropper.paths.getFullFileName())

	stdout, outErr := cmd.CombinedOutput()

	if outErr != nil {
		return errors.New(string(stdout))
	}
	return nil
}
