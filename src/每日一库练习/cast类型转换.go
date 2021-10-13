package main

import (
	"fmt"
	"github.com/spf13/cast"
)

func main() {

	fmt.Println(cast.ToString(1.23))
	fmt.Println(cast.ToString("121321"))
	fmt.Println(cast.ToString("12132t1"))

	fmt.Println("=======")
	fmt.Println(cast.ToInt("21"))
	toInt := cast.ToInt("uuu12")
	fmt.Println(toInt)

	e, err := cast.ToIntE("uu12")
	fmt.Println(e, "==", err)

	//sname := `{"name": "JACK","age":18}`
	//stringE, err := cast.ToStringMapStringE(sname)
	//fmt.Println(stringE, err)

	s := cast.ToString(18219831)
	fmt.Println(s)
}
