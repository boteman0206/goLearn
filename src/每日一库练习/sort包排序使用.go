package main

import (
	"fmt"
	kit "github.com/tricobbler/rp-kit"
	"sort"
)

func main() {

	var ints = []int{1, 56, 2, 6567, 3, 5, 67}

	fmt.Println("排序前：", ints)
	sort.Slice(ints, func(i, j int) bool {
		return ints[i] > ints[j]
	})
	fmt.Println("排序后: ", ints)

	type TES01 struct {
		Code int `json:"code"`
		Sort int `json:"sort"`
	}

	var data []TES01
	for i := 0; i < 10; i++ {
		data = append(data, TES01{
			Code: i,
			Sort: 10 - i,
		})
	}
	fmt.Println("排序前： ", kit.JsonEncode(data))

	sort.Slice(data, func(i, j int) bool {

		return data[i].Code > data[j].Code
	})
	fmt.Println("code排序：", kit.JsonEncode(data))

	sort.Slice(data, func(i, j int) bool {

		return data[i].Sort > data[j].Sort
	})
	fmt.Println("sort排序后：", kit.JsonEncode(data))

}
