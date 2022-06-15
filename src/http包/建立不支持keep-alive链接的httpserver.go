package main

import (
	"log"
	"net/http"
)

/**
假设我们有这样的一个需求，server端完全不支持keep-alive的连接，无论client端发送的请求header中是否显式带有Connection: keep-alive，server端都会在返回应答后关闭连接。那么在Go中，我们如何来实现这一需求呢？我们来看下面代码：

*/
func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("receive a request from:", r.RemoteAddr, r.Header)
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/", Index)
	var s = http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(Index),
	}
	s.SetKeepAlivesEnabled(false) // todo 建立一个不支持keep-alive连接的http server
	s.ListenAndServe()
}
