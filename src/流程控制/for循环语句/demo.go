package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println("hello world！")
	}

	//for ; ;  {
	//	fmt.Println("无线循环1")
	//}

	//for {
	//	fmt.Println("无线循环2")
	//}

	/**
	for-range遍历
	如果字符串中含有中文，那么会出现乱码，传统的是按照字节来取的，可以转成切片来取
	*/
	var addr string = "hello中国上海松江区"
	/**
	for i:=0; i<len(addr); i++ {
		fmt.Println("str: ", string(addr[i])) // 直接遍历会出现乱码
	}
	*/

	str := []rune(addr) // 转换成切片进行遍历
	for i := 0; i < len(str); i++ {
		fmt.Println("str: ", string(str[i]))
	}

	// 使用range遍历
	for key, value := range addr {
		fmt.Println("key : ", key, "  value : ", string(value))
	}

}
