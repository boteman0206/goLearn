package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name  string   `json:"name"`
	Hoppy []string `json:"hoppy"`
	addr  string   `json:"addr"`
}

func main() {

	a := make(map[string]interface{})

	a["name"] = "jack"
	a["hpppy"] = []string{"say", "song"}

	marshal, err := json.Marshal(a)
	fmt.Println(err, " mashal : ", marshal)

	b := make(map[string]interface{})
	err = json.Unmarshal(marshal, &b)
	fmt.Println("err : ", err, "this is a : ", a, "this is b : ", b)
	b["hpppy"] = []string{"baskball", "football"}

	fmt.Println("err : ", err, "this is a : ", a, "this is b : ", b)

	people1 := People{Name: "lucy", Hoppy: []string{"lt", "ny"}, addr: "shanghai"}

	bytes, err := json.Marshal(people1)
	fmt.Println("people1 : ", bytes, err, string(bytes))

	p2 := people1 //  todo 这种拷贝的方式在结构体有slice类型的时候，无法深拷贝
	p2.Name = "pwww333"
	p2.Hoppy[0] = "change"
	fmt.Println(people1, " p2 : ", p2)

	p3 := People{}
	err = json.Unmarshal(bytes, &p3) // todo 使用这种方式深拷贝
	fmt.Println("p3: ", p3, err)
	p3.Hoppy[0] = "深度复制的拷贝"

	fmt.Println("导出的字段类型不包含addr小写 ： ", p3, people1)

}
