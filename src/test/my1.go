package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {

	var sl1 []int
	fmt.Println(sl1)

	sl1 = make([]int, 0)
	fmt.Println(sl1)

	var s1 []int
	var s2 = []int{}
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	fmt.Printf("%#v\n", s1)
	fmt.Printf("%#v\n", s2)

	fmt.Println(len(s1), len(s2))

	//var sa interface{}

	//of := reflect.ValueOf(sa)
	//fmt.Println(of.IsNil())

	var buf bytes.Buffer

	buf.WriteString("hello")
	buf.WriteString("world")
	buf.WriteString("世界")
	buf.WriteString("美好")

	fmt.Println(buf.String())
	buf.Reset()

	data := buf.String()
	fmt.Println("data:", data)

	fmt.Println(strings.EqualFold("你A", "你a"))

	reader := strings.NewReader("hello你好")
	//num, _, _ := reader.ReadRune()
	//fmt.Println(reader)
	//ioutil.WriteFile()
	file, _ := os.OpenFile("string.txt", os.O_CREATE|os.O_APPEND, 777)
	reader.WriteTo(file)

}
