package main

import (
	"fmt"
	"time"
)

func main() {

	te()
}

func te() {
	r := make([]int, 3, 20)
	r[0] = 888

	fmt.Println("888: ", r, len(r), cap(r))
	arr(r)
	fmt.Println("arr after: ", r, len(r), cap(r))
	fmt.Printf("%p\n", r)
}

func arr(r []int) {

	fmt.Printf("%p\n", r)
	fmt.Println("arr1: ", r, len(r), cap(r))
	time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
		r = append(r, i)
		fmt.Printf("%p\n", r)
	}
	r[0] = 999
	fmt.Println("arr2: ", r, len(r), cap(r))
	fmt.Printf("%p\n", r)
}
