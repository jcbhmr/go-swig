//go:build embed

package swigwinzip

import (
	_ "embed"
)

//go:generate curl -fsSL http://prdownloads.sourceforge.net/swig/swigwin-4.2.1.zip -o swigwin.zip

//go:embed swigwin.zip
var swigwinZip []byte

func Get() []byte {
	return swigwinZip
}
