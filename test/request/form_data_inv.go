package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

func main() {

	newTest()
	newTest2()
	handlerTest()
}

func handlerTest() {
	s := http.Server{
		Addr:        ":5000",
		Handler:     nil,
		ReadTimeout: time.Second,
	}
	http.HandleFunc("/", AdvancedHandler)
	log.Fatal(s.ListenAndServe())
}

func AdvancedHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	fmt.Printf("body: %s\n", string(body))

	w.WriteHeader(http.StatusAccepted)
	w.Header().Add("User-Agent", "Mozilla/5.0 (X11; Linux i686; rv:2.0.1) Gecko/20100101 Firefox/4.0.1")
	_, _ = w.Write([]byte("My name is..."))
}

func newTest() {
	formData := url.Values{}
	formData.Add("password", "password")
	formData.Add("role", "mentor")
	formData.Add("username", "userName")

	req, err := http.NewRequest("POST", "http://test.local/api.php", strings.NewReader(formData.Encode()))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("accept", "application/json")
	b, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func newTest2() {
	data := url.Values{}
	data.Set("username", "foo")
	data.Set("surname", "bar")

	r, err := http.NewRequest(http.MethodPost, "http://localhost:5001/advanced", strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		panic(err)
	}
	r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	b, err := httputil.DumpRequestOut(r, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	//c := http.Client{
	//	Timeout: time.Second * 10,
	//}
	//resp, err := c.Do(r)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(resp.Status)
}
