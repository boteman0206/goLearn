package main

import (
	"fmt"
	"reflect"
)

type Coder struct {
	Name string
}

func (c Coder) String() string {
	return c.Name
}

func main() {
	coder := Coder{
		Name: "second coder",
	}

	val := reflect.ValueOf(coder)
	c, ok := val.Interface().(Coder)

	/**
	todo 注意点
	coder := &Coder{
		Name: "second coder",
	}
	c, ok := val.Interface().(*Coder)
	*/

	fmt.Println("coder: ", c, "  ", c.Name, " is ok: ", ok)
	if ok {
		fmt.Println("ok")
	} else {
		panic("error ....")
	}

}
