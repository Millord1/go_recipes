package pdftools

import "strconv"

type PdfToImport struct {
	InFileName  string
	OutFileName string
	InputPath   string
	OutputPath  string
}

func (pdf PdfToImport) CropFile(left int16, top int16, right int16, bottom int16) error {
	// the args are margins, positive to add margin and negative to remove margin
	cropper := newCropper(left, top, right, bottom, pdf)
	return cropper.crop()
}

func (pdf PdfToImport) SplitFile(XDecimation uint16, YDecimation uint16) error {
	splitter := newSplitter(XDecimation, YDecimation, pdf)
	return splitter.split()
}

func intToString(num int16) string {
	return strconv.FormatInt(int64(num), 10)
}
