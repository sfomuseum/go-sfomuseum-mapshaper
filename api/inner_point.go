package api

import (
	"github.com/sfomuseum/go-sfomuseum-mapshaper"
	"net/http"
	"os"
)

func InnerPointHandler(ms *mapshaper.Mapshaper) (http.Handler, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		switch req.Method {
		case "POST":
			// pass
		default:
			http.Error(rsp, "Method not allowed", http.StatusMethodNotAllowed)
		}

		opts := &uploadOptions{
			MaxBytes: 1 * 1024 * 1024,
		}

		tmp_fh, err := uploadWithRequest(rsp, req, opts)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}
		
		defer os.Remove(tmp_fh.Name())

		ctx := req.Context()

		args := []string{
			"-i",
			tmp_fh.Name(),
			"-points",
			"inner",
			"-o",
			"-",
		}

		out, err := ms.Call(ctx, args...)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		rsp.Header().Set("Content-type", "application/json")

		rsp.Write(out)
		return
	}

	h := http.HandlerFunc(fn)
	return h, nil
}
