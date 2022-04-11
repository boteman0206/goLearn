package main

import (
	"fmt"
	"math"
)

func main() {

	var arr []int
	arr = append(arr, 1, 2, 3)
	newSlice := append(arr, 4)
	arr = append(arr, 5)
	arr = append(arr, 6)
	fmt.Println(newSlice)
	fmt.Println(arr)

	even := math.RoundToEven(6.9121)
	fmt.Println("even: ", even)

}
