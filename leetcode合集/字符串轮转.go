package main

import (
	"bytes"
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
	if s1 == "" && s2 == "" {
		return true
	}

	//var data []string
	str := bytes.Buffer{}
	for i := range s2 {
		u := string(s2[i])
		str.WriteString(u)

		fmt.Println("str : ", str.String())

		index := strings.Index(s1, str.String())
		if index == -1 {
			str.Reset()
			str.WriteString(u)

			if i == len(s2)-1 {
				index := strings.Index(s1, str.String())
				if index == -1 {
					return false
				} else {
					return true
				}
			}

		} else {

			if i == len(s2)-1 {
				index := strings.Index(s1, str.String())
				if index == -1 {
					return false
				} else {
					return true
				}
			}
		}

	}

	return false
}

func main() {

	//"waterbottle"
	//"erbottlewat"
	var s1 string = "abcd"

	var s2 string = "acdb"

	//flipedString := isFlipedString(s1, s2)
	flipedString1 := isFlipedString1(s1, s2)

	//fmt.Println("flipedString: ", flipedString)
	fmt.Println("flipedString: ", flipedString1)

}
