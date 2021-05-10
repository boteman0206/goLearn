package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

//todo 相当大的注意点 这里传递的是指针，下面只能使用结构体的地址进行赋值
func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	//var peo People = Student{} //  todo 此时不能将结构体的值赋值给接口类型
	var peo People = &Student{} // todo 只能赋值地址给接口变量
	think := "sb"
	fmt.Println(peo.Speak(think))
}
