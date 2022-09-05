package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func MyTest(r http.ResponseWriter, w *http.Request) {

	//获取post请求参数
	param := &struct {
		Username string `json:"username"`
	}{}

	json.NewDecoder(w.Body).Decode(param)
	fmt.Println("获取post请求参数读取值： ", param)

	//获取get请求参数
	query := w.URL.Query()
	if len(query) > 0 {
		if age, ok := query["age"]; ok {
			fmt.Println("获取get请求参数age: ", age)
		}
		if name, ok := query["name"]; ok {
			fmt.Println("获取get请求参数name: ", name)
		}
	}

	r.Write([]byte("hello world"))

	//fmt.Fprint(r, "Hello world")
}

func main() {

	http.HandleFunc("/hello", MyTest)

	log.Fatal(http.ListenAndServe(":9001", nil))
}
