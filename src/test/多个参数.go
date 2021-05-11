package main

import "fmt"

func Test(name ...string) {

	for i, v := range name {
		fmt.Println(i, " == ", v)
	}

}

func main() {
	Test("jack", "bob", "lucy")
}
