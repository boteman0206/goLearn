package main

import (
	"fmt"
	"reflect"
)

type CoderTest struct {
	Name  string `tag:"myTag"`
	Age   int
	Addr  string
	Hoppy []string
	Mappr map[string]string
}

func main() {
	test := CoderTest{
		Name:  "laseName",
		Age:   12,
		Addr:  "上海",
		Hoppy: []string{"篮球", "桌球"},
		Mappr: map[string]string{
			"one":   "key1",
			"two":   "key2",
			"three": "key3",
		},
	}

	val := reflect.ValueOf(test)
	typ := reflect.TypeOf(test)
	/*
		之前使用，test := &CoderTest{}指针传递出错 of.NumField()必须是指针
		todo panic: reflect: call of reflect.Value.NumField on ptr Value

	*/
	field := val.NumField()

	for i := 0; i < field; i++ {
		fmt.Println("获取的欸一个key ： ", val.Field(i))
		fmt.Println("typ 获取tag ： ", typ.Field(i).Tag.Get("tag"))

	}

}
