package main

import "fmt"

func main() {

	a := 90

	sprint := fmt.Sprint("hello world!", "dsads")
	fmt.Println(sprint)
	sprintf := fmt.Sprintf("%T", a)
	fmt.Println(sprintf)

	//%v 和 %+v的区别
	//%v按值的本来值输出
	//%+v	在 %v 基础上，对结构体字段名和值进行展开

	type Student struct {
		Name string
		Age  int
	}

	student := Student{
		Name: "jack",
		Age:  12,
	}

	fmt.Println(fmt.Sprintf("%v\n", student))  // {jack 12}
	fmt.Println(fmt.Sprintf("%+v\n", student)) // {Name:jack Age:12} 带上了结构体的名称

}
