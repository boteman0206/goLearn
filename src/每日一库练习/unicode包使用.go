package main

import (
	"fmt"
	"unicode"
)

/**
    http://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter02/02.5.html
	utf-8  我们只要记住最后的结果是英文占一个字节，中文占三个字节
	UTF-16 表示最少用两个字节能表示一个字符的编码实现。同样是对 unicode 编码进行转换，它的结果是英文占用两个字节，中文占用两个或者四个字节。
	go 对 unicode 的支持包含三个包 :
		unicode
		unicode/utf8
		unicode/utf16
	unicode 包包含基本的字符判断函数。utf8 包主要负责 rune 和 byte 之间的转换。utf16 包负责 rune 和 uint16 数组之间的转换。
*/

func main() {

	digit := unicode.IsDigit('a')
	fmt.Println("digit: ", digit)

	upper := unicode.IsUpper('你')
	upper1 := unicode.IsUpper('H')
	fmt.Println(upper, upper1)

}
