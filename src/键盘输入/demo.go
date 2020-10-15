package main

import "fmt"

/**
	1: fmt.Scanln 无法读取空格之后的内容，需要使用bufio读取
 */
func main() {

	var name string;
	var addr string;

	//方式一
	fmt.Println("亲输入姓名 ： ")
	fmt.Scanln(&name)
	fmt.Println("name : " , name)

	//方式二
	fmt.Println("请输入地址: ")
	fmt.Scanf("%s", &addr)
	fmt.Println(addr)


}
