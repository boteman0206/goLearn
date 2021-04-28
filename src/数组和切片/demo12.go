package main

import "fmt"

func testFunc(arre [4]int) {
	for key, value := range arre {
		fmt.Println("key: ", key, " val :", value)
	}
}

func main() {
	a1 := []int{1, 2, 3, 4, 5, 455}

	a2 := a1[2:]

	a31 := [...]int{12, 3, 4, 3}
	fmt.Printf("yyy%T\n", a31)
	testFunc(a31)
	fmt.Println(len(a31), cap(a31))

	var a3 = make([]int, len(a1), len(a1))

	copy(a3, a1)

	fmt.Println("a1, a2 a3", a1, a2, a3)
	a1[0] = 100

	fmt.Println("a1, a2 a3", a1, a2, a3)

	ints := append(a2, a3[1:3]...)
	fmt.Println(ints)
	fmt.Println(a1)

}
