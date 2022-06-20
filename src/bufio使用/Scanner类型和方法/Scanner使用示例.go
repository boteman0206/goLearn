package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
我们经常会有这样的需求：读取文件中的数据，一次读取一行。
在学习了 Reader 类型，我们可以使用它的 ReadBytes 或 ReadString来实现，甚至使用 ReadLine 来实现。然而，
在 Go1.1 中，我们可以使用 Scanner 来做这件事，而且更简单好用。

*/

var (
	fileName = "D:\\RpPet\\gitProject\\goLearn\\src\\bytes.buffer\\a.txt"
)

func main() {

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("http://studygolang.com.\nIt is the home of gophers.\nIf you are studying golang, welcome you!")
	// 将文件 offset 设置到文件开头
	file.Seek(0, os.SEEK_SET)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
