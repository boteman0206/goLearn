package main

import (
	"fmt"
	"reflect"
)

func main() {

	var sl1 []int
	fmt.Println(sl1)

	sl1 = make([]int, 0)
	fmt.Println(sl1)

	var s1 []int
	var s2 = []int{}
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	fmt.Printf("%#v\n", s1)
	fmt.Printf("%#v\n", s2)

	fmt.Println(len(s1), len(s2))

	var sa interface{}

	of := reflect.ValueOf(sa)
	fmt.Println(of.IsNil())
}
