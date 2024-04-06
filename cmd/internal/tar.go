package internal

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ExtractTarGzBytesWithStripComponents(tarGzBytes []byte, dest string, stripComponents int) error {
	tarGzStream := bytes.NewReader(tarGzBytes)
	tarStream, err := gzip.NewReader(tarGzStream)
	if err != nil {
		return err
	}
	tarReader := tar.NewReader(tarStream)
	var header *tar.Header
	for header, err = tarReader.Next(); err == nil; header, err = tarReader.Next() {
		// "a/b/c" -> "b/c" -> "c" (using filepath)
		strippedHeaderName := header.Name
		for i := 0; i < stripComponents; i++ {
			strippedHeaderName = filepath.Join(filepath.SplitList(strippedHeaderName)[1:]...)
		}
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(filepath.Join(dest, strippedHeaderName), 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			file, err := os.OpenFile(filepath.Join(dest, strippedHeaderName), os.O_CREATE|os.O_RDWR, header.FileInfo().Mode())
			if err != nil {
				return err
			}
			if _, err := io.Copy(file, tarReader); err != nil {
				_ = file.Close()
				return err
			}
			if err := file.Close(); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unknown type: %v", header.Typeflag)
		}
	}
	if err != io.EOF {
		return err
	}
	return nil
}