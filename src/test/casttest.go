package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"

	"strings"
)

func main() {
	s := cast.ToString(65) + "1"
	fmt.Println("s: ", s)
	trim := strings.Trim(" hello ", " ")

	fmt.Println(len(trim), trim)

	i, i2 := cast.ToInt32E(-3)
	fmt.Println(i, i2)

	i3 := make(map[string]string, 0)
	i3["city"] = "上海"
	i3["address"] = "徐汇区"
	i3["contry"] = "中国"

	bytes, _ := json.Marshal(i3)
	fmt.Println(string(bytes))
	fmt.Println(bytes)
}
