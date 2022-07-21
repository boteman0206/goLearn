package main

import (
	"fmt"
	"reflect"
)

func main() {

	str := "hello"
	num := 8.9

	// 方式一
	fmt.Printf("类型 %T\n", str)
	//fmt.Println(printf, err)

	fmt.Printf("类型 %T\n", num)

	// 方式二
	of := reflect.TypeOf(str)
	fmt.Println("str : ", of)
	typeOf := reflect.TypeOf(num)
	fmt.Println("num：", typeOf)

	//方式三

	var data interface{}
	data = "Hello world"

	switch data.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	default:
		fmt.Println("other")
	}

}
