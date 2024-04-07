//go:build !embed

package swigtargz

import (
	"fmt"
	"net/http"

	"github.com/jcbhmr/go-swig/internal/untar"
)

const Version = "4.2.1"

func ExtractTo(dest string) (err error) {
	url := "http://downloads.sourceforge.net/project/swig/swig/swig-4.2.1/swig-4.2.1.tar.gz"
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download %s: %w", url, err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%s not OK: %s", res.Request.URL, res.Status)
	}
	if res.Header.Get("Content-Type") != "application/x-gzip" {
		return fmt.Errorf("unexpected content type: %s", res.Header.Get("Content-Type"))
	}
	err = untar.Untar(res.Body, dest)
	if err != nil {
		return fmt.Errorf("failed to extract %s to %s: %w", url, dest, err)
	}
	return nil
}
