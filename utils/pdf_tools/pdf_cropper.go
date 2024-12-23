package pdftools

import (
	"os/exec"
	"strings"
)

type PdfCropper struct {
	cmd          string
	arg          string
	MarginLeft   int16
	MarginTop    int16
	MarginRight  int16
	MarginBottom int16
	paths        PdfToImport
}

func NewCropper(left int16, top int16, right int16, bottom int16, pdfPaths PdfToImport) PdfCropper {
	return PdfCropper{
		cmd:          "pdfcrop",
		arg:          "--margins",
		MarginLeft:   left,
		MarginTop:    top,
		MarginRight:  right,
		MarginBottom: bottom,
		paths:        pdfPaths,
	}
}

func (cropper PdfCropper) crop() error {
	var sb strings.Builder
	sb.WriteString(intToString(cropper.MarginLeft))
	sb.WriteString(intToString(cropper.MarginTop))
	sb.WriteString(intToString(cropper.MarginRight))
	sb.WriteString(intToString(cropper.MarginBottom))

	cmd := exec.Command(cropper.cmd, cropper.cmd, sb.String(),
		cropper.paths.InFileName, cropper.paths.OutFileName)

	return cmd.Run()
}
