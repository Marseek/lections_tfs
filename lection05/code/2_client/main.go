package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

func main() {
	v := url.Values{}
	v.Add("id", "1")
	queryString := v.Encode()

	//body := bytes.NewBuffer([]byte("Hello and welcome!"))

	req, err := http.NewRequest(http.MethodPost, "http://google.com/robots.txt"+"?"+queryString, strings.NewReader(v.Encode()))
	if err != nil {
		panic(err)
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux i686; rv:2.0.1) Gecko/20100101 Firefox/4.0.1")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	b, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	c := http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}

	b, err = httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
