package swigwinzp

import (
	"fmt"
	"io"
	"net/http"
)

const Version = "4.2.1"
const url = "http://downloads.sourceforge.net/project/swig/swigwin/swigwin-4.2.1/swigwin-4.2.1.zip"

var data []byte

func Get() ([]byte, error) {
	if data == nil {
		res, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("failed to download %s: %w", url, err)
		}
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("%s not OK: %s", res.Request.URL, res.Status)
		}
		if res.Header.Get("Content-Type") != "application/octet-stream" {
			return nil, fmt.Errorf("unexpected content type: %s", res.Header.Get("Content-Type"))
		}
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read body of %s: %w", res.Request.URL, err)
		}
	}
	return data, nil
}


