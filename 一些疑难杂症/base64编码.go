package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	var str = "ZHONGGUOnihao123"
	strbytes := []byte(str)
	encoded := base64.StdEncoding.EncodeToString(strbytes)

	decoded, err := base64.StdEncoding.DecodeString(encoded)

	decodestr := string(decoded)
	fmt.Println(decodestr, err)

	ciphertext := "eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ"

	// todo 報錯 illegal base64 data at input byte 74
	//todo 大致意思是，如果编码的时候字节不足会在最后加一到两个=号，但看我们的字符串最后没有=，解码的时候解到最后又报错了，赶紧给字符串手动加个等号试试

	//ciphertext = strings.Replace(ciphertext, " ", "", -1)
	ciphertext = fmt.Sprint(ciphertext, "==") // 有些需要有些不需要
	decoded1, err1 := base64.StdEncoding.DecodeString(ciphertext)
	fmt.Println(string(decoded1), err1)

	decoded2, err2 := base64.StdEncoding.DecodeString("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")

	fmt.Println(string(decoded2), err2)

	// todo 解決方式二：直接使用RawURLEncoding 不用追加等号
	encry := "eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ"
	bytes, err := base64.RawURLEncoding.DecodeString(encry)
	fmt.Println(string(bytes), err)
}
