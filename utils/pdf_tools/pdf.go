package pdftools

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
)

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

func (pdf PdfToImport) ReadFile(ingPages []uint16, stepPages []uint16, ingToSkip []string, stepToSkipe []string) error {
	reader := newReader(ingPages, stepPages, ingToSkip, stepToSkipe, pdf)
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

func GetAbsPath(path string) (string, error) {
	// get absolute path
	var err error
	if !filepath.IsAbs(path) {
		path, err = filepath.Abs(path)
	}

	if err != nil {
		return "", err
	}
	return path, nil
}

func (pdf PdfToImport) createDir(dirName string) (string, error) {

	dirPath, err := GetAbsPath(dirName)
	if err != nil {
		logger.Sugar.Panic(err)
		return "", err
	}

	exists, err := dirExists(dirPath)
	if err != nil {
		logger.Sugar.Panic(err)
		return "", err
	}

	if !exists {
		err := os.Mkdir(dirPath, os.ModePerm)
		if err != nil {
			logger.Sugar.Panic(err)
			return "", err
		}
	}

	return dirPath, nil
}

func dirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	return false, err
}
