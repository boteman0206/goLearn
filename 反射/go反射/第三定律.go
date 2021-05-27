package main

import (
	"fmt"
	"reflect"
)

type Coder struct {
	Name string
}

func main() {

	coder1 := &Coder{Name: "three code"}

	of := reflect.ValueOf(coder1)
	// todo 使用val.Elem()，必须要是地址指针， 否则出错
	of.Elem().Field(0).SetString("change code")

	fmt.Println("设置coder  ", coder1)
}
