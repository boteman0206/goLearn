package main

import "fmt"

/**
1：string底层是一个byte数组，因此string也是可以进行切片处理的
2：string是不可变的，不能通过string[0]='z'这种方式修改
3：如果需要修改字符串，可以先将string->[]byte数组或者[]rune， 使用rune可以兼容中文

*/
func main() {

	str := "helloworld!"
	slice := str[3:]
	fmt.Println("slice :", slice)

	//修改string
	arr1 := []byte(str)
	arr1[0] = 't'
	strnew := string(arr1)
	fmt.Println("修改之后：", strnew)

	//修改中文，转成rune
	arr2 := []rune(str)
	arr2[0] = '北'
	strnew1 := string(arr2)
	fmt.Println("修改中文 :", strnew1)

}
