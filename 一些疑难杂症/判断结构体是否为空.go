package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type A struct {
	name string
	age  int
}

func (a A) IsEmpty() bool {
	return reflect.DeepEqual(a, &A{})
}

func main() {
	var a A

	if a == (A{}) { // 括号不能去
		fmt.Println("a == A{} empty")
	}

	if a.IsEmpty() {
		fmt.Println("reflect deep is empty")
	}

	var b *A
	fmt.Println(b == nil)

	i, e := strconv.Atoi("1")
	fmt.Println(i, e)

}
