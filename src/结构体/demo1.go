package main

import (
	"fmt"
)

type Cat struct {
	name  string
	age   int
	color string
	arr1  [5]int

	slice1 []int             // 对于slice是nil没有分配空间需要先make
	m1     map[string]string // 对于map也是nil需要先make

}

func main() {
	var cat Cat
	cat.name = "小白"
	cat.age = 89
	cat.color = "白色"

	fmt.Println(cat)

	//s := Cat{"小黑", 12, "黑色"}
	//fmt.Println("s :", s, s.name, s.age)

	// todo 使用slice和map先make
	cat.m1 = make(map[string]string)
	cat.m1["name"] = "jack"
	cat.slice1 = make([]int, 2)
	cat.slice1[0] = 9

	fmt.Println("cat map ", cat)

}
