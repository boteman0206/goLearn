package main

import (
	"fmt"
	"unsafe"
)

/**
1：bool 类型只能是true或者false
2：默认值是 false
3：占用字节大小 1个字节
*/

func main() {

	var flag bool

	fmt.Println("默认值：", flag)

	//flag = 1 //  定义错误不能用其他来表示bool
	fmt.Println("占用字节大小 ", unsafe.Sizeof(flag))

}
