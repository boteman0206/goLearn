package main

import (
	"fmt"
)

/**
注意点：1: map申明是不会分配内存的，初始化需要make，分配内存之后才能使用
	   2: 无顺序的

*/

func main() {

	var a map[string]string // 没有分配空间
	fmt.Println("a : ", a)
	//a["name"] = "松江"  // 报错,需要先make分配空间

	a = make(map[string]string, 10)
	a["name"] = "松江"
	fmt.Println("make之后： ", a)

	//map的使用方式三种
	//1：先申明在make
	//2：直接make
	var map1 = make(map[string]string)
	fmt.Println("map1 : ", map1)
	//3：直接复制初始化
	var map2 = map[string]string{"name": "jack", "addr": "上海"}
	fmt.Println("map2 : ", map2)

	//delete删除函数, 删除不存在的key不会报错
	delete(map2, "name")
	fmt.Println("delete : ", map2)
	delete(map2, "没有的key")

	//todo：删除所有的key， 1：遍历key逐渐删除 2：make新的map

	//查找map  // 返回两个值有的其中一个是bool， 没有的就是false+空字符
	s, t := map2["name"]
	fmt.Println(s, " t: ", t)
	s1, t1 := map2["addr"]
	fmt.Println(s1, " t1: ", t1)

	//遍历使用for-range遍历
	for k, v := range map2 {
		fmt.Println("k :", k, " v: ", v)
	}

	// len函数
	fmt.Println("len: ", len(map2))

	// map的切片动态的map
	var map3 []map[string]string
	map3 = make([]map[string]string, 2)
	if map3[0] == nil {
		map3[0] = make(map[string]string, 2)
		map3[0]["name"] = "jack"
		map3[0]["addr"] = "上海"
	}

	if map3[1] == nil {
		map3[1] = make(map[string]string, 2)
		map3[1]["name"] = "lucy"
		map3[1]["addr"] = "北京"
	}

	fmt.Println(map3)

	//append函数扩容map
	newMap := make(map[string]string)
	newMap["name"] = "bob"
	newMap["addr"] = "南京"

	strings := append(map3, newMap)
	fmt.Println("newMap :", strings)

}
