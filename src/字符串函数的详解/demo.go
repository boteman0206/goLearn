package main

import (
	"fmt"
	"strconv"
)

func main() {

	// len 按照字节的输出 utf8编码字母一个字节汉子3个字节
	name := "benjing上海"
	fmt.Println("name 的长度 ：", len(name))

	i, err := strconv.ParseInt("12", 10, 64)
	fmt.Println(i, " error : ", err)
}
