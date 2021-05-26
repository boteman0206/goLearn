package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

/**
	todo 错误方式

func main() {
  body := readFile(path)
  fmt.Println(body)
}
func readFile(path string) string {
  f, err := os.Open(path)
  if err != nil {
    log.Fatal(err)
  }
  b, err := ioutil.ReadAll(f)
  if err != nil {
    log.Fatal(err)
  }
  return string(b)
}

*/

/**
todo Go程序使用os.Exit 或者 log.Fatal* 立即退出 (使用panic不是退出程序的好方法，请 don't panic.)
	仅在main（）**中调用其中一个 os.Exit 或者 log.Fatal*。所有其他函数应将错误返回到信号失败中。
*/

// todo 正确的使用方式
func main() {
	body, err := readFile("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(body)
}
func readFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
