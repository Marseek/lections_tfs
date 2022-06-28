package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var data map[string][]*int

type Request_str struct {
	Key  *string `json:"key"`
	Data []*int  `json:"data"`
}

func main2() {

	data = make(map[string][]*int)
	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		request_body := Request_str{}

		err := json.NewDecoder(r.Body).Decode(&request_body)
		if err != nil {
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte(err.Error()))
		}
		for dat := range request_body.Data {
			fmt.Println(dat)
		}
		fmt.Println(request_body)
		for _, data_member := range request_body.Data {
			data[*request_body.Key] = append(data[*request_body.Key], data_member)
		}
		fmt.Println(data[*request_body.Key])
		w.WriteHeader(http.StatusAccepted)
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		s := r.URL.Query().Get("key")
		fmt.Println(s)
		response_data := data[s]

		json.NewEncoder(w).Encode(response_data)
	})

	http.ListenAndServe(":8080", nil)
}

func main1() {

	data = make(map[string][]*int)
	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		request_body := struct {
			Key  *string `json:"key"`
			Data []int   `json:"data"`
		}{}

		err := json.NewDecoder(r.Body).Decode(&request_body)
		if err != nil {
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte(err.Error()))
		}

		for _, data_member := range request_body.Data {
			var a = data_member
			data[*request_body.Key] = append(data[*request_body.Key], &a)
		}

	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		response_data := data[r.URL.Query().Get("key")]
		json.NewEncoder(w).Encode(response_data)
	})
	http.ListenAndServe(":8080", nil)
}

func doSomething(arg int) {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		go func() {
			wg.Add(1)
			// doJob()
			wg.Done()
		}()
	}

}

type Req_str struct {
	Key  string `json:"key"`
	Data int    `json:"data"`
}

func main5() {
	rq := Req_str{Key: "lsdkf", Data: 23}
	bt := jsonWrite(rq)
	go func() {
		panic("ldkf")
	}()
	time.Sleep(time.Second)
	fmt.Println(string(bt))
}

func jsonWrite(data interface{}) []byte {
	raw, _ := json.Marshal(data)
	return raw
}

type myType int

const (
	Sunday = myType(iota)
	Monday
	Tuesday
)

func main() {
	fmt.Println(Sunday, Monday, Tuesday)
	mp := map[string]string{}
	delete(mp, "sel")

}

func iter(data interface{}) []byte {
	raw, _ := json.Marshal(data)
	return raw
}
