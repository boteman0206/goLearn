package main

import (
	"fmt"
)

func main() {

	m := make(map[int][]int, 0)
	m[7] = append(m[7], 5, 2, 3, 4)
	m[8] = append(m[8], 6, 2, 3, 4)
	m[9] = append(m[9], 7, 2, 3, 4)

	for k, v := range m {
		fmt.Println(k, "=======", v)
	}

	for i := range m {
		fmt.Println(i)
	}

}
