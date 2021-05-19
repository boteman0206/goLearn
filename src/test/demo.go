package main

import (
	"fmt"
	"strings"
)

func main() {

	split := strings.Split(":11011", ":")
	fmt.Println("split : ", split)
	for i, v := range split {
		fmt.Println("i : ", i, " v : ", v)
	}
}
