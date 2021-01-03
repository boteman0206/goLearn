package main

import (
	"fmt"
	"sort"
)

func main() {

	s1 := []int{1, 2, 3, 4, 5, 455}

	s2 := s1

	fmt.Println("s1 : s2", s1, s2)

	a1 := append(s1, 99, 88)

	fmt.Println(s1, s2)
	fmt.Println("a1: ", a1)

	s3 := s1[1:3]
	fmt.Println(s3)
	fmt.Printf("%p\n", s3)
	ints := append(s3, 1, 2, 3, 2, 42, 4, 4, 234, 242, 4)
	fmt.Println("ints :", ints, s3)
	fmt.Printf("%p\n, %p\n", ints, s3)

	sort.Ints(ints)

	fmt.Println(ints)
}
