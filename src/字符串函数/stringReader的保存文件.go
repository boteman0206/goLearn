package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println(strings.EqualFold("你A", "你a"))

	reader := strings.NewReader("hello你好")
	//num, _, _ := reader.ReadRune()
	//fmt.Println(reader)
	//ioutil.WriteFile()
	file, _ := os.OpenFile("string.txt", os.O_CREATE|os.O_APPEND, 777)
	reader.WriteTo(file)

}
