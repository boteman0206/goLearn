package main

/**
  func (b *Reader) Peek(n int) ([]byte, error)

同上面介绍的 ReadSlice一样，返回的 []byte 只是 buffer 中的引用，在下次IO操作后会无效，
可见该方法（以及ReadSlice这样的，返回buffer引用的方法）对多 goroutine 是不安全的，也就是在多并发环境下，不能依赖其结果。
*/

import (
	"bufio"
	"fmt"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com.\t It is the home of gophers"), 14)
	go Peek(reader)
	go reader.ReadBytes('\t')
	time.Sleep(1e8)
}

func Peek(reader *bufio.Reader) {
	line, _ := reader.Peek(14)
	fmt.Printf("%s\n", line)
	//time.Sleep(1)
	fmt.Printf("%s\n", line)
}
