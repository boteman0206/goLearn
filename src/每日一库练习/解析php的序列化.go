package main

import "github.com/techoner/gophp"
import "fmt"

/**
todo 专门用于解析php的序列化str，这种php的serize和json的格式不一样
*/

func main() {

	str := `a:1:{s:3:"php";s:24:"世界上最好的语言";}`

	// unserialize() in php
	out, _ := gophp.Unserialize([]byte(str))

	fmt.Println(out) //map[php:世界上最好的语言]

	// serialize() in php
	jsonbyte, _ := gophp.Serialize(out)

	fmt.Println(string(jsonbyte)) // a:1:{s:3:"php";s:24:"世界上最好的语言";}

}
