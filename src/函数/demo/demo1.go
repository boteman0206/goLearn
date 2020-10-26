package demo2

import "fmt"

func init() {
	fmt.Println("导入的包运行init。。。。")
}
func Cal(a int, b int) int {
	return a + b
}

func MultiNum(a, b int) (int, int) {
	return a + b, a - b
}
