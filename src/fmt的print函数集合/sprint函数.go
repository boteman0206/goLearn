package main

import "fmt"

func main() {

	a := 90

	sprint := fmt.Sprint("hello world!", "dsads")
	fmt.Println(sprint)
	sprintf := fmt.Sprintf("%T", a)
	fmt.Println(sprintf)
}
