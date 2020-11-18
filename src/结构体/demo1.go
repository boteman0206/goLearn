package main

import (
	"fmt"
)

type Cat struct {
	name  string
	age   int
	color string
	arr1  [5]int

	slice1 []int             // 对于slice是nil没有分配空间需要先make
	m1     map[string]string // 对于map也是nil需要先make

}

func main() {

	//创建方式一
	var cat Cat
	cat.name = "小白"
	cat.age = 89
	cat.color = "白色"

	fmt.Println(cat)

	//s := Cat{"小黑", 12, "黑色"}
	//fmt.Println("s :", s, s.name, s.age)

	// todo 使用slice和map先make
	cat.m1 = make(map[string]string)
	cat.m1["name"] = "jack"
	cat.slice1 = make([]int, 2)
	cat.slice1[0] = 9

	fmt.Println("cat map ", cat)

	//创建方式二 有多余的字段需要指定属性即可
	s := Cat{name: "小黑", age: 12, color: "黑色"}
	fmt.Println("s : ", s)

	// 方式三 指针方式
	var p3 *Cat = new(Cat)
	fmt.Println("p3: ", *p3)
	(*p3).name = "simth" // 标准写法
	p3.age = 90          // 这是对上面取地址的简化，底层自动执行
	fmt.Println("地址赋值： ", *p3)

	// 方式四 &地址符号
	var p2 *Cat = &Cat{}
	(*p2).name = "lucy"
	p2.age = 88
	fmt.Println("p2 : ", p2)

	var s2 Cat = s // 结构体是进行值拷贝的 ，修改s2的name不会影响s的值

	s2.name = "修改s2"
	fmt.Println("s1修改后 :", s)
	fmt.Println("s2 :", s2)

}
