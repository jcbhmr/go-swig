package internal

import (
	"log"
	"os"

	"github.com/jcbhmr/go-swig/cmd/internal/swigwinzip"
)

func InstallSwig(dest string, l *log.Logger) error {
	if l == nil {
		l = log.Default()
	}
	err := os.MkdirAll(dest, 0755)
	if err != nil {
		return err
	}
	err = ExtractZipBytesWithStripComponents(swigwinzip.Get(), dest, 1)
	if err != nil {
		return err
	}
	return nil
}