package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os"
)

//GBK -> UTF-8
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	all, err := ioutil.ReadAll(reader)
	if err != nil {
		return all, err
	}
	return all, nil
}

//UTF-8 -> GBK
func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	all, err := ioutil.ReadAll(reader)
	if err != nil {
		return all, err
	}
	return all, nil
}

func main() {
	path := `D:\A\桌面\a.txt`

	open, _ := os.Open(path)
	all, _ := ioutil.ReadAll(open)
	utf8, _ := GbkToUtf8(all)
	fmt.Println(string(utf8))
}
