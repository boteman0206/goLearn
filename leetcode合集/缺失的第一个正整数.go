package main

import (
	"fmt"
)

//给定一个未排序的整数数组nums，请你找出其中没有出现的最小的正整数

/**
思路使用map来存储数据，然后从1开始判断有没有存在的key
*/
func minNumberDisappeared(nums []int) int {
	// write code here

	m := make(map[int]struct{}, len(nums))

	for i := range nums {
		m[nums[i]] = struct{}{}
	}

	var key = 1
	for true {
		if _, ok := m[key]; !ok {
			return key
		}
		key++
	}

	return 0
}

func main() {

	var numbers = []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	solution := minNumberDisappeared(numbers)
	fmt.Println(solution)
}
