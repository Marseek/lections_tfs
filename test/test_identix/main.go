package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

const (
	HeaderUserAgent   = "User-Agent"
	HeaderAccept      = "Accept"
	HeaderContentType = "Content-Type"
	UserAgent         = "face_recognizer-http-client/1.0"
	TokenPath         = "/v1/auth/token/"
	ContentTypeJSON   = "application/json"
)

type HTTPClient struct {
	url      string
	userName string
	password string
}

func main() {
	c := HTTPClient{
		url:      "http://identix.kuber.staging.video.rt.ru:80",
		userName: "facerecognizer",
		password: "facerecognizer",
	}

	u, err := url.Parse(c.url)
	if err != nil {
		panic(err)
	}

	u.Path = TokenPath
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	fw, err := writer.CreateFormField("username")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fw, strings.NewReader(c.userName))
	if err != nil {
		panic(err)
	}
	fw, err = writer.CreateFormField("password")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fw, strings.NewReader(c.password))
	if err != nil {
		panic(err)
	}
	err = writer.Close()
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, u.String(), body)
	if err != nil {
		panic(err)
	}

	req.Header.Set(HeaderContentType, writer.FormDataContentType())
	req.Header.Set(HeaderAccept, ContentTypeJSON)
	req.Header.Set(HeaderUserAgent, UserAgent)

	b, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	client := http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	b, err = httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
