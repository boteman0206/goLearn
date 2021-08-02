package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
https://studygolang.com/topics/12072
*/

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandomString(n int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func GetRandomString2(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}
func main() {
	s := GetRandomString(7)
	fmt.Println(s)

	string2 := GetRandomString2(7)
	fmt.Println(string2)
}
