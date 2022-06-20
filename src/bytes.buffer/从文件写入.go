package main

import (
	"bytes"
	"fmt"
	"os"
)

var (
	fileName = "D:\\RpPet\\gitProject\\goLearn\\src\\bytes.buffer\\a.txt"
)

func main() {

	//1: 从文件写入
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Println(file.Sync())
	buf := bytes.NewBufferString("hello ")
	buf.ReadFrom(file) //将text.txt内容追加到缓冲器的尾部
	fmt.Println(buf.String())

	//2: 写出数据到io.Writer  os.File就是实现io.Writer
	fileout, _ := os.OpenFile(fileName, os.O_APPEND, 777)
	buf1 := bytes.NewBufferString("我写入了一个文件")
	buf1.WriteTo(fileout) // hello写到text.txt文件中了

}
