package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	// "time"
)

func toFullname(names ...string) string {
	str := ""
	str1 := ""
	if len(names) > 1 {
		str = names[0]
		str1 = names[1]
	}

	fmt.Println("str: %s", str, "str1:%s", str1)
	fmt.Println("%+v", names)
	return ""
}

func main() {
	toFullname("jack", "pop")
	toFullname()

	// for i := 0; i < 100; i++ {
	// fmt.Println("====", rand.Intn(2))
	// }
	// getFloat()

	fmt.Println(VersionCompare("7.6", "7.6.0"))
	fmt.Println(VersionCompare("7.6.0", "7.6"))
	fmt.Println(VersionCompare("7.6.1", "7.6"))

	// rand.Intn(2)

	fmt.Println("---", 60*10)
	fmt.Println(1703067600 - 60*10)

	a := 90
	if a > 10 || a < 0 {
		fmt.Println("=====", a)
	}

	t := chunkStringSlice([]string{"1", "w", "2", "e", "t", "i", "0"}, 3)

	fmt.Print(t)
}

func getFloat() {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成 0 到 1 之间的随机小数
	randFloat := rand.Float64()

	fmt.Println(randFloat)
	fmt.Println(48 * 60 * 60)
	fmt.Println("------", time.Now().Unix()-1702438822)
}

/**
 * @brief 字符串型版本号比较
 *
 * version format: 1.1.1
 * @return (-1: version1 < version2; 0 v1 = v2 ; 1 v1 > v2 )
 */
func VersionCompare(version1 string, version2 string) int {
	arrVersion1 := strings.Split(version1, ".")
	arrVersion2 := strings.Split(version2, ".")
	lenVer1 := len(arrVersion1)
	lenVer2 := len(arrVersion2)
	lenVer := lenVer1
	if lenVer1 < lenVer2 {
		lenVer = lenVer2
	}
	for i := 0; i < lenVer; i++ {
		var intV1, intV2 int
		if i < lenVer1 {
			intV1, _ = strconv.Atoi(arrVersion1[i])
		}
		if i < lenVer2 {
			intV2, _ = strconv.Atoi(arrVersion2[i])
		}
		if intV1 < intV2 {
			return -1
		}
		if intV1 > intV2 {
			return 1
		}
	}
	return 0
}

func chunkStringSlice(input []string, chunkSize int) [][]string {
	var divided [][]string

	for i := 0; i < len(input); i += chunkSize {
		end := i + chunkSize

		if end > len(input) {
			end = len(input)
		}

		divided = append(divided, input[i:end])
	}

	return divided
}
