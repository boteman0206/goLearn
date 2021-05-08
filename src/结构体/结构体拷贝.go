package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
)

type Dog struct {
	name  string
	color string
	age   int8
	kind  string
}

//todo 带有数组的结构体
type FamilyMember struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Parents []string
}

func main() {
	//1、struct是值类型，默认的复制就是深拷贝
	d1 := Dog{"豆豆", "黑色", 2, "二哈"}
	fmt.Printf("d1: %T , %v , %p \n", d1, d1, &d1)
	d2 := d1 //值拷贝
	fmt.Printf("d2: %T , %v , %p \n", d2, d2, &d2)

	d2.name = "毛毛"
	fmt.Println("d2修改后：", d2)
	fmt.Println("d1：", d1)
	fmt.Println("------------------")

	//2、直接赋值指针地址
	d3 := &d1
	fmt.Printf("d3: %T , %v , %p \n", d3, d3, d3)
	d3.name = "球球"
	d3.color = "白色"
	d3.kind = "萨摩耶"
	fmt.Println("d3修改后：", d3)
	fmt.Println("d1：", d1)
	fmt.Println("------------------")

	fmt.Println("------------------------------")

	s1 := FamilyMember{
		Name:    "jack",
		Age:     12,
		Parents: []string{"dog", "cat"},
	}

	s2 := s1
	var s22 FamilyMember
	// todo 方式一 使用copier
	copier.CopyWithOption(&s22, s2, copier.Option{DeepCopy: true})
	fmt.Println("copier 深度拷贝： ", s22)
	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)
	fmt.Println("修改前。。。。")
	// todo 方式二 使用json
	marshal, err := json.Marshal(s1)
	fmt.Println("err : is : ", err)
	var s23 FamilyMember
	json.Unmarshal(marshal, &s23)
	fmt.Println("s23 de zhi ", s23)

	s2.Parents[0] = "pop"

	fmt.Println("数组修改后的值")
	fmt.Println("s1 : ", s1)
	fmt.Println("s2 : ", s2)
	fmt.Println("s22 : ", s22)
	fmt.Println("s23 de zhi ", s23)

}
