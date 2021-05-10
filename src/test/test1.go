package main

import (
	"fmt"
	"github.com/limitedlee/microservice/common/config"
	"os"
)

var m1 map[string]string

func main() {
	// todo  获取环境变量
	getenv := os.Getenv("ASPNETCORE_ENVIRONMENT")
	fmt.Println("获取环境： ", getenv)
	javaHome := os.Getenv("JAVA_HOME")
	fmt.Println("javaHome： ", javaHome)

	// todo 这个config其实是获取的appsetting.toml文件
	get, err := config.Get("data")
	fmt.Println("getString : ", get, " err ", err)

	m1 = make(map[string]string)
	m1["name"] = "jack"
	fmt.Println(m1)
}
