package pdftools

import "strconv"

type PdfToImport struct {
	FileName  string
	Extension string
	Path      string
}

func (pdf PdfToImport) getPathAndName() string {
	return pdf.Path + pdf.FileName
}

func (pdf PdfToImport) getFullFileName() string {
	return pdf.Path + pdf.FileName + pdf.Extension
}

func (pdf PdfToImport) ReadFile() error {
	reader := newReader(pdf)
	return reader.read()
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
