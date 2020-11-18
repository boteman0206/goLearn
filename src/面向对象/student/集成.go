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
