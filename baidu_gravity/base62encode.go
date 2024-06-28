package main

import (
	"fmt"
	"math"
)

/**
**
** 适用于生成分享码，以及长链接转换短链接的一些服务
 */

// 主要做base62编码操作
var base = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

// 编码
func Base62encode(num int) string {
	baseStr := ""
	for {
		if num <= 0 {
			break
		}

		i := num % 62
		baseStr += base[i]
		num = (num - i) / 62
	}
	return baseStr
}

// 解码
func Base62decode(base62 string) int {
	rs := 0
	len := len(base62)
	f := flip(base)
	for i := 0; i < len; i++ {
		rs += f[string(base62[i])] * int(math.Pow(62, float64(i)))
	}
	return rs
}

func flip(s []string) map[string]int {
	f := make(map[string]int)
	for index, value := range s {
		f[value] = index
	}
	return f
}

func main() {

	str := Base62encode(188)
	fmt.Println("str: ", str)

	num := Base62decode(str)
	fmt.Println("num: ", num)

}
