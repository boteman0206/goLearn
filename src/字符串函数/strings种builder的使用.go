package main

import (
	"fmt"
	"strings"
)

func main() {

	var builder strings.Builder

	name := "hello world!你"
	builder.WriteString(name)

	fmt.Println(builder.String())
	fmt.Println(builder.Len())

	// 重置builder string
	builder.Reset()
	fmt.Println(builder.String())

}
