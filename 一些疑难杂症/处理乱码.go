package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// 转换8进制utf-8字符串到中文
// eg: `\346\200\241` -> 怡
func convertOctonaryUtf8(in string) string {
	s := []byte(in)
	reg := regexp.MustCompile(`\\[0-7]{3}`)

	out := reg.ReplaceAllFunc(s,
		func(b []byte) []byte {
			i, _ := strconv.ParseInt(string(b[1:]), 8, 0)
			return []byte{byte(i)}
		})
	return string(out)
}

func main() {
	s1 := "\346\200\241" // 字面量
	s2 := `\346\200\241` // 原始字符串

	fmt.Println("s1 =", s1)
	fmt.Println("s2 =", s2)

	// 转化 s2
	s3 := convertOctonaryUtf8(s2)
	fmt.Println("s3 =", s3)

	data := `
		{\"msg\":\"\346\210\220\345\212\237\",\"result\":{\"orgCode\":null,\"resultCode\":\"1009001\",\"outSkuId\":null,\"failedDetail\":\"\347\261\273\347\233\256\344\277\256\346\224\271\345\244\261\350\264\245\357\274\214\346\255\243\347\241\256\347\261\273\347\233\256Id\357\274\23224113\357\274\214\347\261\273\347\233\256\345\220\215\347\247\260\357\274\232\345\251\264\345\204\277\347\272\270\345\267\276/\345\251\264\345\204\277\346\271\277\345\267\276\",\"skuId\":null,\"upcCode\":null,\"resultMsg\":\"\345\244\261\350\264\245,\345\216\237\345\233\240\347\274\226\345\217\267[1007085]\"},\"code\":\"0\",\"success\":true,\"detail\":null}
	`

	data = `\350\256\242\345\215\225\345\267\262\347\273\217\345\256\214\346\210\220\357\274\214\345\225\206\345\256\266\344\270\215\345\205\201\350\256\270\345\217\226\346\266\210\350\256\242\345\215\225`
	s4 := convertOctonaryUtf8(data)

	fmt.Println(s4)
}
