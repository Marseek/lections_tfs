package lecture07

import (
	"io"
	"io/ioutil"
	"net/http"
)

func HTTPReq(addr string) (string, error) {
	resp, err := http.DefaultClient.Get(addr)
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	if string(body) == "pong" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, _ = w.Write(body)
	w.WriteHeader(http.StatusOK)
}
