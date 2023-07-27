package main

import (
	"fmt"
	"math/rand"
	"time"
)

type DF struct {
	Name string
}

func main() {

	rand.Seed(time.Now().UnixNano())

	// 0-99
	for i := 0; i < 100; i++ {
		num := rand.Intn(10)
		fmt.Println(num)
	}

	var df DF
	fmt.Println(df.Name)

}
