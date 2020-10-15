package main

import "fmt"

func main() {

	var age int
	fmt.Println("请输入你的年龄 ： ")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("年龄大于18岁")
	} else {
		fmt.Println("年龄太小了！")
	}

}
