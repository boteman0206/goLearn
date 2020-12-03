package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	// 带缓存的文件读取
	fileName := "E:/goProject/goLearn/src/file/a.txt"
	file, e := os.Open(fileName)

	fmt.Println("e : ", e)
	reader := bufio.NewReader(file)

	for {
		str, err1 := reader.ReadString('\n')
		if err1 == io.EOF { // io.EOF表示读取到文件的末尾
			break
		}

		fmt.Println("str : ", str)
	}

	// 一次将文件读取到内存
	bytes, i := ioutil.ReadFile(fileName)
	fmt.Println(string(bytes), " i : ", i)

}
