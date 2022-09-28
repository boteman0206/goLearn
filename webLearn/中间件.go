package main

import (
	"log"
	"net/http"
	"time"
)

func echo(wr http.ResponseWriter, r *http.Request) {
	writeLen, err := wr.Write([]byte("msg hello world! "))
	if err != nil {
		log.Println(err, "write len:", writeLen)
	}

}

func timeMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		timeStart := time.Now()

		next.ServeHTTP(writer, request)

		timeEnd := time.Since(timeStart)
		log.Println("时间差： ", timeEnd)
	})
}

func main() {

	http.Handle("/", timeMiddleware(http.HandlerFunc(echo)))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
