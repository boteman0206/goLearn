package main

import "fmt"

/**
string的使用
1： 默认值是 ""
2:  go中的字符串是不可变的，不能修改
3: 使用反引号输出原生字符串
4： 字符串的拼接  +号实现, 换行+要留在上面
*/
func main() {

	var name string = "jack"
	fmt.Println(name)

	fmt.Println(name[0], string(name[0])) // 可以输出但是不能赋值改变
	//name[0] = 'p' 报错不能更改

	fmt.Println(string(123)) // 输出的是对应的码值

	var price string = "123"
	fmt.Println("price: ", price)

	// 反引号原生字符串输出， 包括换行和特殊字符都不识别， 可以防止攻击

	str3 := `name\nhello""world`
	fmt.Println("反引号输出： ", str3)

	var str4 = "hello" + "world" // +拼接
	//换行+要留在上面
	fmt.Println("str4: ", str4)

}
