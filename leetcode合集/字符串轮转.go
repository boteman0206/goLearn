package main

import (
	"fmt"
	"strings"
)

/*
	字符串轮转。给定两个字符串s1和s2，请编写代码检查s2是否为s1旋转而成（比如，waterbottle是erbottlewat旋转后的字符串）。
	https://leetcode.cn/problems/string-rotation-lcci/

 	输入：s1 = "waterbottle", s2 = "erbottlewat"
*/

// 暴力破解时间比较久
func isFlipedString(s1 string, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}
	if s1 == "" && s2 == "" {
		return true
	}

	var data []string
	for i := range s1 {
		data = append(data, string(s1[i]))
	}

	lenS1 := len(s1)
	for i := 0; i < lenS1; i++ {
		//一个一个旋转去比较，看是否和s2相等
		var spliData []string
		i2 := data[i+1:]
		i3 := data[:i+1]
		spliData = append(spliData, i2...)
		spliData = append(spliData, i3...)
		if strings.Join(spliData, "") == s2 {
			return true
		}
	}

	return false
}

// 采用子串来实现

func isFlipedString1(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 && s1 == "" {
		return true
	}
	repeat := strings.Repeat(s2, 2)
	fmt.Println("repeat: ", repeat)

	index := strings.Index(strings.Repeat(s2, 2), s1)
	if index >= 0 {
		return true
	}
	return false
}

func main() {

	//"waterbottle"
	//"erbottlewat"
	var s1 string = "waterbottle"

	var s2 string = "erbottlewat"

	flipedString := isFlipedString(s1, s2)
	flipedString1 := isFlipedString1(s1, s2)

	fmt.Println("flipedString: ", flipedString)
	fmt.Println("flipedString1: ", flipedString1)

}
