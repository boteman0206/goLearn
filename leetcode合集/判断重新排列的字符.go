package main

import "fmt"

//https://leetcode.cn/problems/check-permutation-lcci/

func CheckPermutation(s1 string, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	m1 := make(map[string]int)
	m2 := make(map[string]int)

	for i := range s1 {
		if num, ok := m1[string(s1[i])]; !ok {
			m1[string(s1[i])] = 1
		} else {
			m1[string(s1[i])] = num + 1
		}
	}

	for i := range s2 {
		if num, ok := m2[string(s2[i])]; !ok {
			m2[string(s2[i])] = 1
		} else {
			m2[string(s2[i])] = num + 1
		}
	}

	for k, v := range m1 {
		key := k
		value := v

		if num, ok := m2[key]; ok {

			if value != num {
				return false
			}
		} else {
			return false
		}

	}

	return true
}

func main() {
	s1 := "abc"
	s2 := "bca"

	permutation := CheckPermutation(s1, s2)
	fmt.Println(permutation)
}
