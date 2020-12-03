package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "E:/goProject/goLearn/src/file/a.txt"

	file, e := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 777)
	fmt.Println(file, e)

	writer := bufio.NewWriter(file)

	for i := 0; i < 5; i++ {
		writer.WriteString("hello 新增！\n")
	}

	writer.Flush()
	file.Close()

}
