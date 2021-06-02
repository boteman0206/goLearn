package main

import "fmt"

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func main() {

	a, b := 2, 3
	max := If(a > b, a, b).(int)
	println(max)

	a1 := "jwt1"
	b1 := "sos"

	b2 := len(a1) > len(b1)
	i := If(b2, a1, b1).(string)

	fmt.Println(" this is string : ", i)

}
