package main

import (
	"fmt"
	"strings"
)

//https://leetcode.cn/problems/string-to-url-lcci/
/**

URL化。编写一种方法，将字符串中的空格全部替换为%20。假定该字符串尾部有足够的空间存放新增字符，
并且知道字符串的“真实”长度。（注：用Java实现的话，请使用字符数组实现，以便直接在数组上操作。）

*/
func replaceSpaces1(S string, length int) string {

	var newStr string

	leth := 0
	for i := range S {
		s := string(S[i])
		if s != " " {
			newStr += s
		} else {
			newStr += "%20"
		}

		leth += 1
		if leth >= length {
			return newStr
		}
	}

	return newStr
}

func replaceSpaces(S string, length int) string {
	build := strings.Builder{}
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			build.WriteString("%20")
		} else {
			build.WriteByte(S[i])
		}
	}
	return build.String()
}

func main() {

	//输入："Mr John Smith    ", 13
	//输出："Mr%20John%20Smith"

	s1 := "Mr%20John Smith    "
	spaces := replaceSpaces(s1, 13)
	println(s1, len(s1), " spaces", spaces)

	//for i := range s1 {
	//	s := string(s1[i])
	//	fmt.Println(s, " 是否是空格：  ", s == " ")
	//}

	runes := []rune(s1)
	r1 := runes[:13]
	fmt.Println(runes, len(runes), r1, len(r1))

}
