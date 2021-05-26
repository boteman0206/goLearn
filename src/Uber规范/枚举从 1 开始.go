package main

import "fmt"

type Operation int

/**
todo 在 Go 中引入枚举的标准方法是声明一个自定义类型和一个使用了 iota 的 const 组。由于变量的默认值为 0，因此通常应以非零值开头枚举。
*/
const (
	Add Operation = iota + 1
	Subtract
	Multiply
)

// Add=1, Subtract=2, Multiply=3

func main() {

	fmt.Println(Add, Subtract, Multiply)
}
