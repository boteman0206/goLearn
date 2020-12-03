package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	src := "C:/Users/Administrator/Desktop/新建文件夹 (2)/PT6A9423.jpg"
	dst := "E:/goProject/goLearn/src/file/b.jpg"

	file, _ := os.Open(src)

	openFile, _ := os.OpenFile(dst, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 777)
	reader := bufio.NewReader(file)

	writer := bufio.NewWriter(openFile)

	written, err := io.Copy(writer, reader)

	fmt.Println(written, " err: ", err)

}
