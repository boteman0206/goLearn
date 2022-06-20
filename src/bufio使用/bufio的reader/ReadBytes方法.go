package main

import (
	"bufio"
	"fmt"
	"strings"
)

/**
ReadBytes和ReadSlice功能和用法都很像，那他们有什么不同呢？

在讲解ReadSlice时说到，它返回的[]byte 是指向 Reader 中的 buffer，而不是 copy 一份返回，也正因为如此，通常我们会使用 ReadBytes 或 ReadString。
很显然，ReadBytes 返回的 []byte 不会是指向 Reader 中的 buffer，通过查看源码可以证实这一点。

*/
func main() {

	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	line, _ := reader.ReadBytes('\n')
	fmt.Printf("the line:%s\n", line)
	// 这里可以换上任意的 bufio 的 Read/Write 操作
	n, _ := reader.ReadBytes('\n')
	fmt.Printf("the line:%s\n", line)

	fmt.Println(string(n))

}
