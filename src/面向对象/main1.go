package main

import "./接口"

func main() {

	// 接口的调用 接口是引用赋值

	phone := 接口.Phone{}
	computer := 接口.Computer{}
	computer.Run(phone)

	/**
	todo 所有的类型都实现了空接口，可以接受任何类型
	*/

}
