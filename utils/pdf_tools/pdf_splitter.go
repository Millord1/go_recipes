package pdftools

import (
	"os/exec"
)

type PdfSplitter struct {
	cmd         string
	arg         string
	XDecimation uint16
	YDecimation uint16
	paths       PdfToImport
}

func NewSplitter(XDecimation uint16, YDecimation uint16, pdfPaths PdfToImport) PdfSplitter {
	return PdfSplitter{
		cmd:         "mutool",
		arg:         "poster",
		XDecimation: XDecimation,
		YDecimation: YDecimation,
		paths:       pdfPaths,
	}
}

func (splitter PdfSplitter) split() error {
	cmd := exec.Command(splitter.cmd, splitter.arg,
		"-x "+intToString(int16(splitter.XDecimation)),
		"-y "+intToString(int16(splitter.YDecimation)),
		splitter.paths.InFileName, splitter.paths.OutFileName)

	return cmd.Run()
}
