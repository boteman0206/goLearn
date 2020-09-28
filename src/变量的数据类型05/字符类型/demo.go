package main

import "fmt"

/**
golang 中没有专门的字符类型， 如果要存储单个字符（字母），
一般使用byte来保存
*/

func main() {

	var c1 byte = 'a'
	var c2 byte = '0'
	//输出使用ascII码
	fmt.Println("c1: ", c1, "c2 : ", c2)
	// 格式化输出
	fmt.Printf("c1=%c  c2=%c", c1, c2)

	//var name byte = '级'  //  overflows byte 报错
	//fmt.Println("中文 ：", name)
	var name int = '级'
	fmt.Printf("name = %c， c3的码值 %d", name, name)

	fmt.Println()
	/**
	遍历字符串
	*/
	for _, name := range "hello" {

		fmt.Println(name, string(name))
	}

	/**
	字符使用的细节
	1： 使用中文的时候保存要用int因为码值已经超过ascii码了
	2： go语言使用utf-8编码
	3： %c使用格式化
	4： 字符类型可以进行运算
	*/

	fmt.Println("字符运算：", 'a'+'b')

}
