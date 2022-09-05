package main

import (
	"fmt"
	"net/http"
)

//https://www.topgoer.com/gin%E6%A1%86%E6%9E%B6/gin%E6%BA%90%E7%A0%81%E8%A7%A3%E8%AF%BB/%E6%8F%AD%E5%BC%80gin%E7%9A%84%E7%A5%9E%E7%A7%98%E9%9D%A2%E7%BA%B1.html
func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	//http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Hello World"))
	//})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("start http server fail:", err)
	}
}

/**

http.HandleFunc没有请求方法类型


*/
