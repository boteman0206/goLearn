package main

import (
	"fmt"
	"github.com/limitedlee/microservice/common/config"
	"os"
)

type Operation int

const (
	Add Operation = iota + 1
	Subtract
	Multiply
	EnvVar = 8
	Testy
	Testy1 = iota
	testy2
)

var m2 = map[string]string{}

var m1 map[string]string

func main() {

	fmt.Println("testy is : ", Testy, Testy1, testy2) // 5
	m2["name"] = "lucy"
	fmt.Println("m2: ", m2)

	s := make([]int, 0)
	fmt.Println("长度为0 ： ", len(s), s)
	s = nil
	s = append(s, 12)
	fmt.Println("len的长度： ", len(s), s)

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

	fmt.Println(int32(11) / int32(2))

	//var a int32
	//a = int32(float32(11)/float32(2))
	//fmt.Println("a: ", a)
}
