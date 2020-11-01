package main

import "fmt"

func errorFunc() {

	// 使用defer + recover 来捕获异常
	defer func() {
		err := recover() // 内置的recover函数，可以捕获异常
		if err != nil {  // 说明捕获到异常
			fmt.Println("err : ", err)
			// 处理错误逻辑的代码，可以关闭流文件或者发送错误日志给管理员
		}
	}()

	num1 := 10
	num2 := 0
	i := num1 / num2
	fmt.Println("i : ", i)
}

func main() {
	/*
		go不适用try...catch和finally
		使用defer, panic, recover

	*/

	errorFunc()
	fmt.Println("捕获异常之后的输出。。。")

}
