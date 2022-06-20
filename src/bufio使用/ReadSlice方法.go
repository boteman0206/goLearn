package main

import (
	"bufio"
	"fmt"
	"strings"
)

/**
ReadSlice 从输入中读取，直到遇到第一个界定符（delim）为止，返回一个指向缓存中字节的 slice，在下次调用读操作（read）时，这些字节会无效。举例说明：

从结果可以看出，第一次ReadSlice的结果（line），在第二次调用读操作后，内容发生了变化。
也就是说，ReadSlice 返回的 []byte 是指向 Reader 中的 buffer ，而不是 copy 一份返回。正因为ReadSlice 返回的数据会被下次的 I/O 操作重写，
因此许多的客户端会选择使用 ReadBytes 或者 ReadString 来代替。
*/
func main() {

	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	line, _ := reader.ReadSlice('\n')
	fmt.Printf("the line:%s\n", line)
	// 这里可以换上任意的 bufio 的 Read/Write 操作
	n, _ := reader.ReadSlice('\n')
	fmt.Printf("the line:%s\n", line) // 这里打印的变化了
	fmt.Println(string(n))

	fmt.Println("====testError===============")
	testError()
}

/**
todo 如果 ReadSlice 在找到界定符之前遇到了 error ，它就会返回缓存中所有的数据和错误本身（经常是 io.EOF）。
如果在找到界定符之前缓存已经满了，ReadSlice 会返回 bufio.ErrBufferFull 错误。
当且仅当返回的结果（line）没有以界定符结束的时候，ReadSlice 返回err != nil，
也就是说，如果ReadSlice 返回的结果 line 不是以界定符 delim 结尾，那么返回的 err也一定不等于 nil（可能是bufio.ErrBufferFull或io.EOF）
*/
func testError() {
	reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com"), 16)

	line, err := reader.ReadSlice('\n')
	fmt.Printf("line:%s\terror:%s\n", line, err)
	line, err = reader.ReadSlice('\n')
	fmt.Printf("line:%s\terror:%s\n", line, err)
}
