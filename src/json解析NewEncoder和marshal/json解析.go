package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type People struct {
	Name string
	Age  int
}

func main() {

	//1：序列化结构体
	people := People{Name: "jack", Age: 12}

	// todo json格式化必须要使用大写
	bytes, e := json.Marshal(people)

	fmt.Println(bytes, " erro : ", e, " string ", string(bytes))

	// 2： 序列化map
	i := make(map[string]interface{})
	i["name"] = "lucy"
	i["age"] = 9

	marshal, i2 := json.Marshal(i)
	fmt.Println("map: ", string(marshal), " i2", i2)

	//反序列化
	var t1 People
	unmarshal := json.Unmarshal(marshal, &t1) // todo 必须使用引用传递，才能改变t1的值
	//unmarsha2 := json.Unmarshal(marshal, t1) // 这样执行没有值  Unmarshal(non-pointer main.People)
	fmt.Println(unmarshal, t1)

	//json转成map
	var map1 map[string]interface{}
	decoder := json.NewDecoder(strings.NewReader(string(marshal))).Decode(&map1)

	fmt.Println(map1)
	fmt.Println("decoder: ", decoder)
}
