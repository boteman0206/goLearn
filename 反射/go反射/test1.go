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

	val := reflect.ValueOf(&test)
	typ := reflect.TypeOf(test)
	/*
		之前使用，test := &CoderTest{}指针传递出错 of.NumField()必须是指针
		todo panic: reflect: call of reflect.Value.NumField on ptr Value

	*/
	field := typ.NumField()

	for i := 0; i < field; i++ {
		fmt.Println("获取的欸一个key ： ", val.Elem().Field(i))
		get := typ.Field(i).Tag.Get("tag")
		fmt.Println("typ 获取tag ： ", get)
		s := val.Elem().Field(i).String()
		if s == "laseName" && get == "myTag" {
			fmt.Println("名称是laseName 并且 tag 是 myTag")
			name := typ.Field(i).Name
			byName := val.Elem().FieldByName(name)
			fmt.Println("filed name is :", byName)

			//todo 修改结构体内容
			set := val.Elem().Field(i).CanSet()
			if set {
				val.Elem().Field(i).SetString("my change name")
			}
		}
	}

	fmt.Println("修改之后的name ： ", test)

}
