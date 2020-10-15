package main

import (
	"fmt"
)

func main() {

	var age int
	fmt.Println("请输入你的年龄 ： ")
	fmt.Scanln(&age)

	if age > 58 { // 不需要加小括号
		fmt.Println("老油条")
	} else if age > 18 {
		fmt.Println("有点东西！")
	} else {
		fmt.Println("太年轻！")
	}

}
