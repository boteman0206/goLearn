package main

import (
	"fmt"
)

/**
给定一个逆波兰表达式，求表达式的值。

数据范围：表达式长度满足 1 \le n \le 10^4 \1≤n≤10
4
   ，表达式中仅包含数字和 + ，- , * , / ，其中数字的大小满足 |val| \le 200 \∣val∣≤200 。

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

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param tokens string字符串一维数组
 * @return int整型
 */
func evalRPN(tokens []string) int {
	// write code here

	pushStack := new(stack)

	m := make(map[string]struct{}, 0)
	var explain = []string{"+", "-", "*", "/"}
	for i := range explain {
		m[explain[i]] = struct{}{}
	}

	for i := range tokens {
		if _, ok := m[tokens[i]]; ok {
			front1 := pushStack.front()
			pushStack.pop()
			var f1 int
			fmt.Sscanf(front1, "%d", &f1)

			var f2 int
			front2 := pushStack.front()
			pushStack.pop()
			fmt.Sscanf(front2, "%d", &f2)

			var sprintf string
			switch tokens[i] {
			case "+":
				i2 := f1 + f2
				sprintf = fmt.Sprintf("%d", i2)
			case "-":
				i2 := f1 - f2
				sprintf = fmt.Sprintf("%d", i2)
			case "*":
				i2 := f1 * f2
				sprintf = fmt.Sprintf("%d", i2)
			case "/":
				i2 := f1 / f2
				sprintf = fmt.Sprintf("%d", i2)
			}

			pushStack.push(sprintf)

		} else {
			pushStack.push(tokens[i])
		}

	}
	var res int
	fmt.Sscanf(pushStack.front(), "%d", &res)
	return res
}

func main() {

	a := []string{"3", "0", "-"}

	rpn := evalRPN(a)
	fmt.Println("rpn :;; ", rpn)

	s := "id:00123"

	var i int
	if _, err := fmt.Sscanf(s, "id:%5d", &i); err == nil {
		fmt.Println(i + 9) // Outputs 123
	}

	s1 := "90"
	_, err := fmt.Sscanf(s1, "%d", &i)
	if err != nil {
		return
	}

	fmt.Println(i)

}
