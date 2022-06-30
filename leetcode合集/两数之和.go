package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	kit "github.com/tricobbler/rp-kit"
)

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

	var data bytes.Buffer
	json.NewEncoder(&data).Encode("{'name':'jack', 'age':19}")

	type user struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var user1 user

	json.NewDecoder(&data).Decode(&user1)
	fmt.Println("this is use1 : ", kit.JsonEncode(user1))

}
