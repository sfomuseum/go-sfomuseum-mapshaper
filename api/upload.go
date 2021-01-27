package api

import (
	// "errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type uploadOptions struct {
	MaxBytes int64
}

func uploadWithRequest(rsp http.ResponseWriter, req *http.Request, opts *uploadOptions) (*os.File, error) {

	req.Body = http.MaxBytesReader(rsp, req.Body, opts.MaxBytes)

	defer req.Body.Close()

	tmp_fh, err := ioutil.TempFile("", "mapshaper-*.geojson")

	if err != nil {
		return nil, err
	}

	_, err = io.Copy(tmp_fh, req.Body)

	if err != nil {
		return nil, err
	}

	return tmp_fh, nil
}
