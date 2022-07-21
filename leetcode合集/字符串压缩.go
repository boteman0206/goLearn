package main

import "fmt"

/**
字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，
字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。
来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/compress-string-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

 输入："aabcccccaaa"
 输出："a2b1c5a3"

*/

func compressString(S string) string {

	if len(S) == 1 {
		return S
	}

	var str string
	var num int = 1
	var one string
	for i := range S {
		one = string(S[i])
		if i > 0 {
			next := string(S[i-1])
			if next == one {
				num++
				if i == len(S)-1 {
					str += fmt.Sprintf("%s%d", next, num)
				}
			} else {

				str += fmt.Sprintf("%s%d", next, num)
				num = 1
				if i == len(S)-1 {
					str += fmt.Sprintf("%s%d", one, num)
				}

			}
		}
	}
	if len(str) > len(S) {
		return S
	}

	return str
}

func main() {

	var str = "aabcccccaa"
	str = "a"
	//for i := range str {
	//fmt.Println(str[i])
	//}

	s := compressString(str)
	fmt.Println(s)
}
