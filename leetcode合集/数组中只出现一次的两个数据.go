package main

import "fmt"

//一个整型数组里除了两个数字只出现一次，其他的数字都出现了两次。请写程序找出这两个只出现一次的数字

func FindNumsAppearOnce(array []int) []int {
	// write code here

	m := make(map[int]struct{}, len(array))
	for i := range array {
		if _, ok := m[array[i]]; ok {
			delete(m, array[i])
		} else {
			m[array[i]] = struct{}{}
		}
	}

	var res []int
	for k, _ := range m {
		res = append(res, k)
	}
	return res
}

func main() {

	var array = []int{1, 4, 1, 6}
	fmt.Println(FindNumsAppearOnce(array))
}
