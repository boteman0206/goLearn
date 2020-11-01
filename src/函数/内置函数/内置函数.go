package 内置函数

import "fmt"

func main() {
	// len函数，字符长度按照字节

	// new函数 用来分配值空间，比如int，float，struct返回的是指针
	// make函数 用来分配内存，主要用来分配引用类型，比如chan，map，slice
	i := new(int)
	fmt.Println(i)

}
