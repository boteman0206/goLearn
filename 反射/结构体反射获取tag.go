package main

import (
	"fmt"
	"reflect"
)

type Persons struct {
	name string `json:"abc"`
	age  int    `from:"gi"`
	addr string `json:"pop"`
}

func main() {

	per := Persons{"jack", 12, "北京"}

	of := reflect.TypeOf(&per)
	fmt.Println("kine : ", of.Elem().Kind(), of.Kind())

	field0 := of.Elem().Field(0) // 指针的获取方式
	fmt.Println("filed0 : ", field0)

	of1 := reflect.TypeOf(per)
	fmt.Println("kine : ", of1.Kind())
	field1 := of1.Field(0) // 类型获取
	fmt.Println("filed : ", field1, field1.Type, field1.Name)

	fmt.Println(field1.Tag.Get("json")) // 获取json标间

	fmt.Println(of1.NumField()) // 结构体有几个字段

	valueOf := reflect.ValueOf(&per)
	fmt.Println(valueOf.Elem().Kind(), valueOf.Elem().NumField())

}
