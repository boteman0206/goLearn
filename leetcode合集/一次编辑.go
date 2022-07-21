package main

import (
	"fmt"
	"strings"
)

/**
字符串有三种编辑操作:插入一个英文字符、删除一个英文字符或者替换一个英文字符。
给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
https://leetcode.cn/problems/one-away-lcci/

输入:
first = "pale"
second = "ple"
输出: True

*/

// 暴力破解
func oneEditAway(first string, second string) bool {

	if first == second && first == "" {
		return true
	}

	var st []string
	var ss []string
	for i := range first {
		st = append(st, string(first[i]))
	}

	for i := range second {
		ss = append(ss, string(second[i]))
	}
	if len(first) == len(second) { // 替换
		//找出不同的字符，进行替换，看替换之后是否一样
		for i := range second {
			var d2 = make([]string, len(st))
			copy(d2, st)
			d2[i] = string(second[i])
			if strings.Join(d2, "") == second {
				return true
			}
		}
	}
	if len(first) > len(second) && len(first)-1 == len(second) { // 删除
		for i := range st {
			var d2 = make([]string, len(st))
			copy(d2, st)
			d2[i] = ""
			if strings.Join(d2, "") == second {
				return true
			}
		}
	}
	if len(first) < len(second) && len(second)-1 == len(first) { // 新增
		for i := range ss {
			var d2 = make([]string, len(ss))
			copy(d2, ss)
			d2[i] = ""
			if strings.Join(d2, "") == first {
				return true
			}
		}
	}

	return false
}

func main() {
	s1 := "pale"
	s2 := "palej"
	away := oneEditAway(s1, s2)

	fmt.Println(away)

	var d1 = []int{1, 2, 34}
	var d2 = make([]int, len(d1))
	i := copy(d2, d1)

	fmt.Println(d2, d1, i)

	index := strings.Index("hello", "l")
	fmt.Println(index)
}
