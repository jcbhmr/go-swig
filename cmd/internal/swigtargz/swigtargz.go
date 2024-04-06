//go:build !embed

package swigtargz

import (
	"io"
	"net/http"
)

var swigwinTarGz []byte

func Get() []byte {
	if swigwinTarGz == nil {
		url := "http://prdownloads.sourceforge.net/swig/swig-4.2.1.tar.gz"
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		swigwinTarGz, err = io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
	}
	return swigwinTarGz
}
