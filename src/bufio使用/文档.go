package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

/**
// bufio 包实现了带缓存的 I/O 操作
bufio 用来帮助处理 I/O 缓存。 我们将通过一些示例来熟悉其为我们提供的：Reader, Writer and Scanner 等一系列功能



*/
func main() {

	// 1: 使用bufio.NewReader构造一个reader
	inputReadBuf := strings.NewReader("1234567890")
	reader := bufio.NewReader(inputReadBuf)

	var buf bytes.Buffer

	reader.WriteTo(&buf)

	fmt.Println(buf.String())

	//2： 使用NewReaderSize构建一个reader
	reader = bufio.NewReaderSize(strings.NewReader("http://studygolang.com"), 16)

	reader.WriteTo(&buf)
	fmt.Println(buf.String())

}
