package main

/**
基于 Go 1.18+ 泛型（map、filter、contains、find...）的 Lodash 风格的 Go 库

Lodash: 是一个javascript风格的库
*/

import (
	"fmt"
	"github.com/samber/lo"
	"reflect"
	"strconv"
)

func main() {
	names := lo.Uniq[string]([]string{"Samuel", "Marc", "Samuel"})
	// []string{"Samuel", "Marc"}
	fmt.Println("names: ", names)

	even := lo.Filter[int]([]int{1, 2, 3, 4}, func(x int, y int) bool {
		fmt.Println(x, y) // 找出奇数的数据 y代表的是index索引
		return x%2 == 0
	})
	fmt.Println(even)

	imap := lo.Map[int64, string]([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
		//return cast.ToString(x)
	})
	fmt.Println(imap, reflect.TypeOf(imap))
	var data string
	for i := range imap {
		data += imap[i]
	}
	fmt.Println(data)

}
