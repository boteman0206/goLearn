package main

import (
	"fmt"
)
import "../demo"     // 引用方式一： 包名demo和文件包的名demo2不一致
import "../test"     // 引用方式二： 包名test和文件包名test一致
import abc "../test" // 别名的方式导入，使用时候用abc去调用

/**
包的快速入门
1： 包名和文件夹的名称通常保持一致
2： 使用首字母大写的函数或者变量才可以导出， 类似于java的public关键字
3： import包时， 路径是从$GOPATH的src下面开始的，
	这个需要配置，可以使用相对路径来解决
*/
func main() {

	cal := demo2.Cal(16, 18) // 使用方式： 包名和文件名不一致
	fmt.Println("cal : ", cal)

	s := test.Test("hello ,", "world!")
	fmt.Println("s : ", s) // 使用方式： 包名和文件名一致

	s1 := abc.Test("hello ,", "world!")
	fmt.Println("别名: ", s1)

	// cd到main包的文件目录下
	//go build -o my.exe 变编译成二进制文件 取别名
	//go build ./main.go 变编译成二进制文件 名称为main.exe

	//返回多个值得在计算方式
	i, i2 := demo2.MultiNum(30, 4)
	fmt.Println("第一个： ", i, " 第二个值： ", i2)
}
