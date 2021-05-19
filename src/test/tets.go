package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	newInt := big.NewInt(123)
	fmt.Println("newInt: ", newInt)

	t, err := time.Parse("2006-01-02", "2020-10-14")
	fmt.Println(t, err)
	t, err = time.Parse("2006-01-02 15:04:05", "2020-10-14 00:00:00")
	fmt.Println(t, err)

	//time.ParseInLocation()
}
