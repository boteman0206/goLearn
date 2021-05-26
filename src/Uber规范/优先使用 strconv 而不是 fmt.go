package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {

	b := []int{1, 23, 4, 5}

	for i := 0; i < len(b); i++ {
		s := fmt.Sprint(rand.Int())
		fmt.Println(s)
	}

	for i := 0; i < len(b); i++ {

		fmt.Println("随机数 int31： ", rand.Int31())
		fmt.Println("随机数 int： ", rand.Int())
		fmt.Println("随机数 intn： ", rand.Intn(10))
		s := strconv.Itoa(rand.Intn(12))

		num, err := strconv.Atoi("1213")
		fmt.Println(s, num, err)
	}
}
