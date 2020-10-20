package main

import "fmt"

/**
1：break默认跳转出最近的for循环
2：break设置跳转标签
*/
func main() {

	//label1: // 设置跳转标签
	for i := 0; i < 4; i++ {
	label2: // 设置跳转标签
		for j := 0; j < 10; j++ {
			if j == 2 {
				break label2
			}
			fmt.Println("j = :", j)
		}
	}

	// continue 可以使用标签跳转
label3: // 设置跳转标签
	for i := 0; i < 4; i++ {
		//label4:  // 设置跳转标签
		for j := 0; j < 10; j++ {
			if j == 2 {
				continue label3
			}
			fmt.Println("j = :", j)
		}
	}

	/**
	goto语句
	*/
	n := 11
	fmt.Println("ok1")
	if n > 10 {
		goto label5 // 直接跳转到label5
	}
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
label5:
	fmt.Println("ok5")
	fmt.Println("ok6")

}
