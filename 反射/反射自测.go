package main

import (
	"fmt"
	"reflect"
	"sync"
)

func testInterface(v interface{}) {
	of := reflect.ValueOf(v)

	of.Type()

	i := of.Interface()
	fmt.Println("i 的类型 ： ", i)
	fmt.Println(fmt.Sprintf("i的类型%T", i))
	fmt.Println(of.Kind())

	switch of.Type().Kind() {
	case reflect.Int:
		fmt.Println("int 类型 。。。")
	}

	switch i.(type) {
	case int:
		fmt.Println("===", "interesting int ")
	}

	once := sync.Once{}
	once.Do(func() {
		fmt.Println("hello world!")
	})
}

func main() {
	var num = 90
	testInterface(num)
}
