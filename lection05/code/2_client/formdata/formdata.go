package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	run()
}

func run() {
	_ = call("http://localhost:8080/employee", "POST")
}

func call(urlPath, method string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fw, _ := writer.CreateFormField("name")
	_, _ = io.Copy(fw, strings.NewReader("John"))

	fw, _ = writer.CreateFormField("age")
	_, _ = io.Copy(fw, strings.NewReader("23"))

	fw, _ = writer.CreateFormFile("photo", "test.png")
	file, _ := os.Open("test.png")
	_, _ = io.Copy(fw, file)

	// Close multipart writer.
	writer.Close()
	req, err := http.NewRequest("POST", "http://localhost:8080/employee", bytes.NewReader(body.Bytes()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rsp, _ := client.Do(req)
	if rsp.StatusCode != http.StatusOK {
		log.Printf("Request failed with response code: %d", rsp.StatusCode)
	}
	return nil
}
