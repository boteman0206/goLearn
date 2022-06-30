package main

import (
	"fmt"
)

func MoreThanHalfNum_Solution(numbers []int) int {
	// write code here
	lenN := len(numbers)

	m := make(map[int]int, len(numbers))

	for i := range numbers {
		if v, ok := m[numbers[i]]; ok {
			m[numbers[i]] = v + 1
		} else {
			m[numbers[i]] = 1
		}
	}

	for k, v := range m {
		if v > lenN/2 {
			return k
		}
	}

	return 0
}

func main() {

	var numbers = []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	solution := MoreThanHalfNum_Solution(numbers)
	fmt.Println(solution)

}
