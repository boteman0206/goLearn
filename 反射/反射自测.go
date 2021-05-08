package main

import (
	"fmt"
	"reflect"
)

func testInterface(v interface{}) {
	of := reflect.ValueOf(v)
	i := of.Interface()
	fmt.Println("i 的类型 ： ", i)
	fmt.Println(of.Kind())

	switch of.Type().Kind() {
	case reflect.Int:
		fmt.Println("int 类型 。。。")
	}

	switch i.(type) {
	case int:
		fmt.Println("===", "interesting int ")
	}

}

func main() {
	var num = 90
	testInterface(num)
}
