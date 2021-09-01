package main

import (
	"bytes"
	"fmt"
	"go.uber.org/zap/buffer"
	"strings"
)

func main() {

	var builder strings.Builder

	name := "hello world!你"
	builder.WriteString(name)

	fmt.Println(builder.String())
	fmt.Println(builder.Len())

	// 重置builder string
	builder.Reset()
	fmt.Println(builder.String())

	var buf buffer.Buffer

	buf.WriteString("shijie")
	buf.WriteString("世界")
	fmt.Println(buf.String())

	i := bytes.Buffer{}
	i.WriteString("kejfpoefwe世界")

	fmt.Println(i.Bytes())

}
