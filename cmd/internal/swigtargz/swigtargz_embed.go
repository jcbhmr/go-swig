//go:build embed

package swigtargz

import (
	_ "embed"
)

//go:generate curl -fsSL http://prdownloads.sourceforge.net/swig/swig-4.2.1.tar.gz -o swig.tar.gz

//go:embed swig.tar.gz
var swigwinTarGz []byte

func Get() []byte {
	return swigwinTarGz
}
