package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

/**
案例一
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is log run ....")
	io.WriteString(w, "[v2 this is] Hello, Kubernetes!")
}


func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":3000", nil)
*/

func hello(w http.ResponseWriter, r *http.Request) {

	host, _ := os.Hostname()
	fmt.Println("存活探针版本 run ....", host)
	io.WriteString(w, fmt.Sprintf("[存活探针版本] Hello, Kubernetes! host: %s", host))
}

func main() {

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("开始探活了  run。。。。", time.Now().Format(time.RFC3339))
		w.Write([]byte("ok[healthz successfull]"))
	})

	http.HandleFunc("/", hello)
	http.ListenAndServe(":3001", nil)
}
