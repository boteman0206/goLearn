package main

import "fmt"

func twoSum(nums []int, target int) []int {

	m := make(map[int]int, len(nums))
	for i := range nums {
		m[nums[i]] = i
	}

	for i := range nums {
		lookUp := target - nums[i]
		if k, ok := m[lookUp]; ok && k != i {
			return []int{i, k}
		}
	}
	return nil
}

func main() {

	var num = []int{3, 2, 4}
	var target = 6
	sum := twoSum(num, target)

	fmt.Println(sum)
}
