package main

import (
	"fmt"
)

//给定一个未排序的整数数组nums，请你找出其中没有出现的最小的正整数

func minNumberDisappeared(nums []int) int {
	// write code here

	if len(nums) == 1 {
		return nums[0]
	}

	return 0
}

func main() {

	var numbers = []int{-1, 2, 3, 2, 2, 2, 5, 4, 2}
	solution := minNumberDisappeared(numbers)
	fmt.Println(solution)
}
