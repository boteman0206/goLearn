package main

import (
	"fmt"
	"sort"
)

/**
排序的map
1： 先将key放到int数组中
2： 排序int数组
3：遍历int数组，输出map
这个貌似只能输出key为数字类型的，其他的排序貌似是不可以的
*/

func main() {

	map1 := make(map[int]string)
	map1[1] = "a"
	map1[3] = "b"
	map1[2] = "c"
	map1[4] = "d"

	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	fmt.Println(keys)
	sort.Ints(keys)
	fmt.Println("keys : ", keys)

	for e, val := range keys {
		fmt.Println(e, " val :", val)
		fmt.Println(map1[val])
	}

}
