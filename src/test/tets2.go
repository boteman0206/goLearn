package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("recovered:")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("not good")
}
