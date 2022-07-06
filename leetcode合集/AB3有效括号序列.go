package main

import "fmt"

/*
给出一个仅包含字符'(',')','{','}','['和']',的字符串，判断给出的字符串是否是合法的括号序列
括号必须以正确的顺序关闭，"()"和"()[]{}"都是合法的括号序列，但"(]"和"([)]"不合法。

数据范围：字符串长度 0\le n \le 100000≤n≤10000
要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)


*/

type stack struct {
	stack []string
}

func (s *stack) push(data string) {
	s.stack = append(s.stack, data)
}

func (s *stack) pop() {
	if len(s.stack) > 0 {
		s.stack = s.stack[:len(s.stack)-1]
	}
}

func (s *stack) len() int {
	return len(s.stack)
}

func (s *stack) front() string {
	if len(s.stack) > 0 {
		return s.stack[len(s.stack)-1]
	}
	return ""
}

func (s *stack) first() string {
	if len(s.stack) > 0 {
		return s.stack[0]
	}
	return ""
}

func (s *stack) firstPop() {
	if len(s.stack) > 0 {
		s.stack = s.stack[1:len(s.stack)]
	}
}

func isValid(s string) bool {
	// write code here

	pushStack := new(stack)

	for i := range s {
		// 左括号入栈
		if string(s[i]) == "(" {
			pushStack.push("(")
		}
		if string(s[i]) == "{" {
			pushStack.push("{")
		}
		if string(s[i]) == "[" {
			pushStack.push("[")
		}

		// 有括号出栈
		if string(s[i]) == ")" {
			front := pushStack.front()
			if front != "(" {
				return false
			}
			pushStack.pop()
		}
		if string(s[i]) == "}" {
			front := pushStack.front()
			if front != "{" {
				return false
			}
			pushStack.pop()
		}
		if string(s[i]) == "]" {
			front := pushStack.front()
			if front != "[" {
				return false
			}
			pushStack.pop()
		}
	}

	// 特殊情况单个符号的处理： [或者{或者(
	if pushStack.len() > 0 {
		return false
	}
	return true
}

func main() {

	a := "["
	valid := isValid(a)
	fmt.Println(valid)

}
