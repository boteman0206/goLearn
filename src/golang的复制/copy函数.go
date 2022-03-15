package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	fmt.Println(slice1, "  clice2: ", slice2)
	slice1[0] = 99
	fmt.Println(slice1, "改变后：", slice2)

	copy(slice1, slice2)

	fmt.Println(slice1, " ", slice2)
	slice2[0] = 90
	fmt.Println(slice1, " ", slice2)

	// todo copy函数先定义要copy到的对象中 var s = make([]int32, 0)，就是深拷贝

	type user struct {
		name string
		age  int32
	}

	var s1 = []user{{
		name: "elk",
		age:  12,
	}, {
		name: "abc",
		age:  10,
	}}

	var s2 = make([]user, len(s1))

	copy(s2, s1)
	fmt.Println(s1, " s2 ", s2)

}
