package main

import (
	"bytes"
	"fmt"
)

/**
bytes.buffer是一个缓冲byte类型的缓冲器，这个缓冲器里存放着都是byte。

*/

//如何创建bytes.buffe
func main() {

	buf1 := bytes.NewBufferString("hello")
	fmt.Println(buf1)
	buf2 := bytes.NewBuffer([]byte("hello"))
	fmt.Println(buf2)
	buf3 := bytes.NewBuffer([]byte{byte('h'), byte('e'), byte('l'), byte('l'), byte('o')})
	fmt.Println(buf3)
	// 以上三者等效

	buf4 := bytes.NewBufferString("")
	fmt.Println(buf4)
	buf5 := bytes.NewBuffer([]byte{})
	fmt.Println(buf5)
	// 以上两者等效

}
