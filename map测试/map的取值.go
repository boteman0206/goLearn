package main

import "fmt"

func main() {

	m := make(map[string]string)
	m["name"] = "jack"

	k, isBool := m["name"]
	fmt.Println(" k ", k, " is bool : ", isBool)

	k1, isBool1 := m["age"]
	fmt.Println(" k ", k1, " is bool : ", isBool1)

}
