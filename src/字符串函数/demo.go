package main

import (
	"fmt"
	"strings"
)

func main() {

	index := strings.Index("hello", "l")
	fmt.Println("index : ", index)

	i := strings.Index("你好北京", "北")
	fmt.Println("i : ", i)

	lastIndex := strings.LastIndex("hello", "l")
	fmt.Println("lastIndex : ", lastIndex)

	// 忽略大小写的比较
	fold := strings.EqualFold("abc", "ABc")
	fmt.Println("fold : ", fold)

	// 替换
	replace := strings.Replace("hello", "l", "p", 2)
	replace1 := strings.Replace("hello", "l", "p", -1) // 全部都换
	fmt.Println("replace : ", replace)
	fmt.Println("replace : ", replace1)

	// 切割
	split := strings.Split("hello,world, ok", ",")
	splitN := strings.SplitN("hel,lo,wo,rld,ok", ",", 3)
	fmt.Println("split : ", split)
	fmt.Println("aplitN : ", splitN)

	lower := strings.ToLower("AASD")
	upper := strings.ToUpper("adsa")
	fmt.Println("LOWER :", lower)
	fmt.Println("upper :", upper)

	space := strings.TrimSpace("  sdas   ")
	fmt.Println("space : ", space, len(space))
	// 去掉左右两边的空格和！号
	trim := strings.Trim("  !ad!as!  ", " !a")
	fmt.Println("trim : ", trim)

	// 以什么开头
	prefix := strings.HasPrefix("hello", "h")
	fmt.Println("prefix : ", prefix)
	suffix := strings.HasSuffix("hello", "O")
	fmt.Println("suffix : ", suffix)

	contains := strings.Contains("hello", "he")
	fmt.Println("contians : ", contains)
	any := strings.ContainsAny("hwllo", "ahda")
	fmt.Println("any : ", any)
}
