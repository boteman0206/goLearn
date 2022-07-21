package main

import "fmt"

//https://leetcode.cn/problems/palindrome-permutation-lcci/

/**

给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。
回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。
回文串不一定是字典当中的单词。

*/

// 只有一个是奇数 其余的都是偶数菜有可能是回文数子
func canPermutePalindrome(s string) bool {

	m := make(map[string]int, 0)

	for i := range s {
		s2 := string(s[i])
		if num, ok := m[s2]; !ok {
			m[s2] = 1
		} else {
			m[s2] = num + 1
		}
	}

	var num int
	for _, v := range m {
		if v%2 != 0 {
			num++
		}
	}

	if num == 1 || num == 0 {
		return true
	} else {
		return false
	}
}

func main() {

	s := "tactca"
	palindrome := canPermutePalindrome(s)

	fmt.Println(palindrome)

}
