package main

import (
	"bytes"
	"fmt"
)

func main() {

	//1： 写入string
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString("小花猫")
	fmt.Println(buf.String())

	readRune, i, err := buf.ReadRune() // 读取一个run数据 小
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	fmt.Println("readRune:", readRune, string(readRune), " i ", i)

	fmt.Println(buf.String()) // 剩余的rune数据  花猫

	fmt.Println("============ 写入[]byte================")
	//2： 写入[]byte
	buf1 := bytes.NewBuffer([]byte{})
	s := []byte("小黑猫")
	buf1.Write(s)
	fmt.Println(buf1.String())

	var b byte = '?'
	buf1.WriteByte(b)

	fmt.Println(buf1.String())
	fmt.Println("====写入rune==================")
	var r = '小'
	buf1.WriteRune(r)
	fmt.Println(buf1.String())

}
