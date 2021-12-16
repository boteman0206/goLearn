package main

import (
	"fmt"
	"reflect"
)

type TetsReflect struct {
	Name string `json:"name"`
	Age  int    `json:"age"`

	SomeOthers string `json:"some_others"`
}

func (this *TetsReflect) ReflectMethod1() {

	fmt.Println("这是方法一：", this.Name, " age :", this.Age, " some ", this.SomeOthers)
}

func (this *TetsReflect) ReflectMethod2() {

	fmt.Println("这是方法二：", this.Name, " age :", this.Age, " some ", this.SomeOthers)
}

func main() {

	var data = &TetsReflect{
		Name:       "方法一",
		Age:        10,
		SomeOthers: "我是方法一",
	}

	ofType := reflect.TypeOf(data)
	kind := ofType.Kind()
	fmt.Println("kind :", kind) // 查看kind的类型

	//ofValue := reflect.ValueOf(data)

	if kind == reflect.Ptr { // todo 如果是指针类型在使用下面的ofType.NumField() 时候要用值，不能用指针

		ofType = ofType.Elem()
		//修改完之后的type
		i := ofType.Kind()
		fmt.Println("修改完之后的type : ", i.String(), " 是否等于true ：", i == reflect.Struct)
	}

	field := ofType.NumField() // 还有是struct的时候才能够用这个方法

	for i := 0; i < field; i++ {
		structField := ofType.Field(i)
		s := structField.Name
		fmt.Println("字段 : ", i, " 内容 ", s)
		jsonTag := structField.Tag.Get("json")
		fmt.Println("tag : ", jsonTag)

	}

	ofType1 := reflect.TypeOf(data) // todo 方法的时候又要使用地址来
	method := ofType1.NumMethod()
	for i := 0; i < method; i++ {
		methodName := ofType1.Method(i)
		fmt.Println(methodName.Name)
	}

}
