package main

import (
	"fmt"
	"reflect"
)

func justifyType(x interface{}) {

	of := reflect.ValueOf(x)
	if of.Kind() == reflect.Ptr {
		/*
			TODO 如果是地址类型的话就转成原始的类型来判断
		*/
		fmt.Println("this is ptr :", of.Kind())
		elem := of.Elem()
		fmt.Println("element : ", elem)
		x = elem.Interface()
	}

	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}

func main() {
	justifyType(18)
	justifyType("123")
	justifyType(true)

	var a = 89

	justifyType(&a)
}
