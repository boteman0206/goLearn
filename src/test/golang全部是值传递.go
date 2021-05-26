package main

import "fmt"

func hello(m map[string]string) {
	m["name"] = "hello 修改" // 修改map
}
func main() {

	m := make(map[string]string)
	fmt.Println("main : ", len(m))

	m["name"] = "jack"

	fmt.Println("之前： ", m)

	hello(m)

	fmt.Println("之后：", m)

}
