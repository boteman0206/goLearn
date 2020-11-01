package main

import (
	"errors"
	"fmt"
)

/**
自定义错误信息
1：errors.New("错误说明")
2： panic内置函数，接受一个接口的类型作为参数，输出错误信息，并且推出程序
*/

func readConfig(name string) error {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("配置文件错误...")
	}
}

func test02() {
	err := readConfig("config.ini112")
	if err != nil {
		// 抛出异常，终止程序
		// 如果读取文件序错误就输出这个错误，并且终止程序
		panic(err)
	}

	fmt.Println("test02()继续执行。。。")
}

func main() {

	test02()
	fmt.Println("main函数执行")

}
