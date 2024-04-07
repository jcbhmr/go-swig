//go:build embed

package swigtargz

import (
	"bytes"
	"fmt"
	_ "embed"

	"github.com/jcbhmr/go-swig/internal/untar"
)

const Version = "4.2.1"

//go:generate curl -fsSL http://prdownloads.sourceforge.net/swig/swig-4.2.1.tar.gz -o swig.tar.gz

//go:embed swig.tar.gz
var swigTarGz []byte

func ExtractTo(dest string) error {
	err := untar.Untar(bytes.NewReader(swigTarGz), dest)
	if err != nil {
		return fmt.Errorf("failed to extract swig-%s (embedded) to %s: %w", Version, dest, err)
	}
	return nil
}
