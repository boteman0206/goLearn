package student

import "fmt"

type Animols struct {
	Name string
	Age  int
}

type Cat struct {
	Animols
}

type Dog struct {
	Animols
}

func (a Animols) Speak() {
	fmt.Println("animous :", a.Name)
}

func (a Animols) Test() {
	fmt.Println("animous 的test方法 :", a.Name)
}

func (p Dog) Test() {
	fmt.Println("dog的专用方法。。。", p.Name)
}
