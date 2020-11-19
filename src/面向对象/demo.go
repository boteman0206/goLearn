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

	// 继承中公用方法使用
	dog.Animols.Speak()
	cat.Animols.Speak()
	cat.Speak()
	dog.Speak()

	dog.Test()         // 会首先使用dog的Test方法，没有的时候在使用Animous的Test方法
	dog.Animols.Test() // 这个会直接使用animous的Test

	//todo 注意事项
	/*
		1: 结构体可以使用匿名结构体中的所有字段和方法（包括大写和小写）
		2：子类的字段会覆盖父类结构体中的字段，先去子类中找对应的字段（就近访问原则）
		3： 如果继承了多个结构体，并且子类中没有指定的字段，必须指定某一个匿名结构体的字段属性，否则会报错
		4：嵌套有名的接固体必须指定结构体的名称
		type D struct {
			a A
		}
		必须带上 d.a.name 才可以

	*/

	i := &student.Dog{student.Animols{"jack", 19}}
	fmt.Println("I : ", i.Animols.Name)
	fmt.Println("I : ", i.Name)

}
