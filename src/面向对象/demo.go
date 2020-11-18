package main

import (
	"./student"
	"fmt"
	"os"
)

type Student struct {
	Name string
	Age  int
}

func main() {

	environ := os.Environ()
	fmt.Println(environ)

	getuid := os.Getuid()
	fmt.Println(getuid)

	student1 := Student{Name: "jack", Age: 19}
	student2 := Student{Name: "tom", Age: 23}
	student3 := Student{Name: "lucy", Age: 27}
	fmt.Println(student1, student1.Name, student1.Age)
	fmt.Println(student2)
	fmt.Println(student3)

	//student := &Student{Name: "pop", Age: 12}
	//fmt.Println(student)

	person := student.NewPerson("akdla", 12, "上海")
	fmt.Println("person : ", person)

	cat := &student.Cat{}
	cat.Animols.Name = "小花猫"
	cat.Animols.Age = 19

	dog := &student.Dog{}
	dog.Animols.Name = "牧羊犬"
	dog.Animols.Age = 12

	// 公用方法使用
	dog.Animols.Speak()
	cat.Animols.Speak()
}
