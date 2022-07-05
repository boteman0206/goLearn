package main

import (
	"bytes"
	"fmt"
)

type stack struct {
	stack []int32
}

func (s *stack) push(data int32) {
	s.stack = append(s.stack, data)
}

func (s *stack) pop() {
	if len(s.stack) > 0 {
		s.stack = s.stack[:len(s.stack)-1]
	}
}

func (s *stack) front() int32 {
	if len(s.stack) > 0 {
		return s.stack[len(s.stack)-1]
	}
	return 0
}

func (s *stack) first() int32 {
	if len(s.stack) > 0 {
		return s.stack[0]
	}
	return 0
}

func (s *stack) firstPop() {
	if len(s.stack) > 0 {
		s.stack = s.stack[1:len(s.stack)]
	}
}

func removeStr(data string) string {
	var st = new(stack)

	for i := range data {
		if int32(data[i]) == st.front() {
			st.pop()
		} else {
			st.push(int32(data[i]))
		}
	}

	var b bytes.Buffer
	for true {
		front := st.first()
		if front == 0 {
			break
		} else {
			b.WriteString(string(front))
			st.firstPop()
		}

	}

	return b.String()
}

var (
	inputData string
)

func main() {

	//str := "abba"
	//fmt.Println(string(str[0]))
	//
	//for i := range str {
	//	fmt.Println(str[i])
	//}
	//
	//s := removeStr(str)
	//fmt.Println("---s-", s)
	//
	//var data = []int{1, 2, 3, 4, 5, 56, 6}
	//
	//fmt.Println(data[:len(data)-1])

	//fmt.Println("Please enter your full name: ")
	fmt.Scanln(&inputData)

	//fmt.Println("inputData: ", inputData)
	str := removeStr(inputData)
	if len(str) == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(str)
	}

}
