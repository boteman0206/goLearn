package main

import (
	"fmt"
	"io"
	"net/http"
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
	fmt.Println("健康检查 run ....")
	io.WriteString(w, "[健康检查] Hello, Kubernetes!")
}

func main() {

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("开始探活了  run。。。。", time.Now().Format(time.RFC3339))
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/", hello)
	http.ListenAndServe(":3000", nil)
}
