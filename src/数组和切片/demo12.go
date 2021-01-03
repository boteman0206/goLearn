package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3, 4, 5, 455}

	a2 := a1[2:]

	var a3 = make([]int, len(a1), len(a1))

	copy(a3, a1)

	fmt.Println("a1, a2 a3", a1, a2, a3)
	a1[0] = 100

	fmt.Println("a1, a2 a3", a1, a2, a3)

	ints := append(a2, a3[1:3]...)
	fmt.Println(ints)
	fmt.Println(a1)

}
