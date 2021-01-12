package main

import (
	"fmt"
	"reflect"
)

func testValue(x interface{}) {

	of := reflect.ValueOf(x)
	fmt.Println("of : ", of)
	kind := of.Kind()

	fmt.Println("kind : ", kind)

	//fmt.Println(of.Elem()) 不是指针类型无法使用

}

/**
指针类型的使用
*/
func testValuePtr(x interface{}) {

	of := reflect.ValueOf(x)
	fmt.Println("of : ", of)
	kind := of.Kind()
	fmt.Println("kind : ", kind)

	fmt.Println(of.Elem())        // todo 获取指针的值
	fmt.Println(of.Elem().Kind()) // todo 指针获取类型

	of.Elem().SetInt(122)

	fmt.Println(of.Elem())
}

func main() {

	a1 := 12
	testValue(a1)
	fmt.Println("============")
	testValuePtr(&a1)

	fmt.Println("main a1 : ", a1)
}
