package internal

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path/filepath"
)

func ExtractZipBytesWithStripComponents(zipBytes []byte, dest string, stripComponents int) error {
	zipStream := bytes.NewReader(zipBytes)
	zipReader, err := zip.NewReader(zipStream, int64(zipStream.Len()))
	if err != nil {
		return err
	}
	for _, zipFile := range zipReader.File {
		if zipFile.FileInfo().IsDir() {
			if err := os.MkdirAll(filepath.Join(dest, zipFile.Name), 0755); err != nil {
				return err
			}
			continue
		}
		zipFileReader, err := zipFile.Open()
		if err != nil {
			return err
		}
		// "a/b/c" -> "b/c" -> "c" (using filepath)
		strippedFileName := zipFile.Name
		for i := 0; i < stripComponents; i++ {
			strippedFileName = filepath.Join(filepath.SplitList(strippedFileName)[1:]...)
		}
		file, err := os.OpenFile(filepath.Join(dest, strippedFileName), os.O_CREATE|os.O_RDWR, zipFile.Mode())
		if err != nil {
			_ = zipFileReader.Close()
			return err
		}
		if _, err := io.Copy(file, zipFileReader); err != nil {
			_ = file.Close()
			_ = zipFileReader.Close()
			return err
		}
		if err := file.Close(); err != nil {
			_ = zipFileReader.Close()
			return err
		}
		if err := zipFileReader.Close(); err != nil {
			return err
		}
	}
	return nil
}