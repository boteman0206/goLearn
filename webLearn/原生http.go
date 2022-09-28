package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func echo(wr http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		fmt.Println("this is get http")
	}
	if r.Method == "POST" {
		fmt.Println("this is post http ")
	}

	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wr.Write([]byte("echo error"))
		return
	}

	writeLen, err := wr.Write(msg)
	if err != nil || writeLen != len(msg) {
		log.Println(err, "write len:", writeLen)
	}

}

/**
因为默认的 net/http 包中的 mux 不支持带参数的路由 不管是get post delete put 都会进入echo方法，可以在里面通过r.Method
所以市面上出现了很多http框架
开源界有这么几种框架，第一种是对 httpRouter 进行简单的封装，然后提供定制的中间件和一些简单的小工具集成比如 gin 主打轻量，易学，高性能
第二种是借鉴其它语言的编程风格的一些 MVC 类框架，例如 beego，方便从其它语言迁移过来的程序员快速上手，快速开发
还有一些框架功能更为强大，除了数据库 schema 设计，大部分代码直接生成，例如 goa
*/

func main() {

	http.HandleFunc("/", echo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
