package main

import "fmt"

func main() {

	s := "hello"
	b := make([]byte, len(s))
	//var name string
	copy(b, s)

	fmt.Println(string(b))

}
