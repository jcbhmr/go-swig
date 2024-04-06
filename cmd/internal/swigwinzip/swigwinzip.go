//go:build !embed

package swigwinzip

import (
	"io"
	"net/http"
)

var swigwinZip []byte

func Get() []byte {
	if swigwinZip == nil {
		url := "http://prdownloads.sourceforge.net/swig/swigwin-4.2.1.zip"
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		swigwinZip, err = io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
	}
	return swigwinZip
}
