package main

import (
	"fmt"
	"reflect"
)

func threeSum(num []int) [][]int {
	// write code here
	if len(num) <= 2 {
		return [][]int{}
	}

	var arr [][]int
	for i := range num {
		for i2 := range num {
			if i == i2 {
				continue
			}
			for i3 := range num {
				if i == i3 || i2 == i3 {
					continue
				}
				var newArr []int
				newArr = append(newArr, num[i])
				newArr = append(newArr, num[i2])
				newArr = append(newArr, num[i3])
				arr = append(arr, newArr)
			}
		}
	}
	fmt.Println(" arr ： ", arr)

	var resArr [][]int
	for i := range arr {
		ints := arr[i]
		var sum int
		for i2 := range ints {
			sum += ints[i2]
		}
		if sum == 0 {
			resArr = append(resArr, ints)
		}

	}
	fmt.Println("resArr: ", resArr)

	//去重

	return nil
}

func main() {

	//[-10,0,10,20,-10,-40]     [[-10,-10,20],[-10,0,10]]
	num := []int{-10, 0, 10, 20, -10, -40}
	threeSum(num)
	// [-2,0,1,1,2]   [[-2,0,2],[-2,1,1]]

	// [0,0]  []
	a := []int{-10, 0, 10}
	b := []int{-10, 0, 10}
	fmt.Println(reflect.DeepEqual(a, b))

	fmt.Println(TestEqual(a, b))

}

func TestEqual(a, b []int) bool {

	var flag = true
	for i := range a {
		if a[i] != b[i] {
			flag = false
		}
	}
	return flag
}
