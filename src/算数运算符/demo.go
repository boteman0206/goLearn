package main

import "fmt"

/**

 */

func main() {

	/**
	除法使用的特点
	*/
	// 除法  都是整数的话运算之后保留整数
	fmt.Println(10 / 4)

	var n1 float32 = 10 / 4
	fmt.Println("n1 : ", n1) // n1 :  2

	//如果需要保留小数则需要保留有float参与运算
	fmt.Println(10.0 / 4)

	/**
	取模的特点  取模的符号由第一位决定
	*/
	fmt.Println(10 % 3)   // 1
	fmt.Println(-10 % 3)  // -1
	fmt.Println(10 % -3)  // 1
	fmt.Println(-10 % -3) // -1

	/**
	++i 和  --i 没有 i++和i-- 只能独立使用
	*/
	var i int = 1
	i++
	fmt.Println(i)

}
