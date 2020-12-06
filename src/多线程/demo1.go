package main

import "fmt"

func test() {
	for i := 0; i < 100; i++ {
		fmt.Println("test 执行 。。。", i)
	}
}

func main() {

	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main 执行。。。。 ", i)
	}
}
