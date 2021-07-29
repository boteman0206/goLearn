package main

import (
	"fmt"
	"github.com/spf13/cast"
	"strings"
)

func main() {
	s := cast.ToString(65) + "1"
	fmt.Println("s: ", s)
	trim := strings.Trim(" hello ", " ")

	fmt.Println(len(trim), trim)

	i, i2 := cast.ToInt32E(-3)
	fmt.Println(i, i2)
}
