package main

import "fmt"

func main() {

	fmt.Println("tom\tjack") // 制表符号  \t

	fmt.Println("git\nhub") // 换行 \n
	fmt.Println("=========")

	fmt.Println("回车被覆盖\r测试") // 回车不换行 \r
	// 测试会覆盖回车两个字 但是在IDE中不会正确的显示， 可以使用go run 进行正确的输出

	fmt.Println("hello \"i love \"world") //  \斜杠进行分割

	//练习
	fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t12\t河北\t北京")

}
