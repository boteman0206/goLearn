package main

import (
	"fmt"
	"regexp"
)

const text = "My email is ccmouse@gmail.com"

func main() {
	compile := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	match := compile.FindString(text)
	fmt.Println("match: ", match)
}
