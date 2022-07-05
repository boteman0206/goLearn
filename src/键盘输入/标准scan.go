package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	f := bufio.NewReader(os.Stdin) //读取输入的内容

	for {
		fmt.Println("请输入一些字符串：")
		readString, err := f.ReadString(',')
		if err != nil {
			return
		}
		if len(readString) <= 0 {
			continue
		}
		fmt.Println("你的输入是", readString)
		var input string
		_, err = fmt.Sscan(readString, &input)
		if err != nil {
			return
		}

		if input == "exit" {
			return
		}

	}
}
