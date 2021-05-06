package main

import (
	"fmt"
	"github.com/unknwon/com"
	"reflect"
)

/**
反射使用的包 reflect
*/

type Student struct {
	name string
	age  int
}

func (s Student) Test() {
	fmt.Println("test run ...", s.name)
}

func reflectMethod(t interface{}) {
	of := reflect.TypeOf(t)
	fmt.Println("of :", of)
	fmt.Println(of.Kind())
	fmt.Println(of.NumField())
	fmt.Println(of.FieldByName("name"))

	fmt.Println(of.NumMethod()) // 大写的方法才能反射
	fmt.Println(of.MethodByName("Test"))

	valueOf := reflect.ValueOf(t)
	fmt.Println("valueOf: ", valueOf)

	// value转成空接口
	i := valueOf.Interface()

	// 空接口类型断言转成原来的结构体
	student, ok := i.(Student)
	fmt.Println("类型断言 ： ", ok, student.name, student.age)
	int1, ok := i.(int)
	fmt.Println("错误的类型断言：", ok, int1)

}

func main() {
	student := Student{name: "jack", age: 12}
	student.Test()

	reflectMethod(student)

	mustInt := com.StrTo("ffggg").MustInt()
	fmt.Println(mustInt)

}
