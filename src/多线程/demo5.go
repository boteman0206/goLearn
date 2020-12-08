package main

import "fmt"

type Cat struct {
	name string
}

func main() {

	// 存放任意数据类型
	var allChan chan interface{}

	allChan = make(chan interface{}, 10)

	allChan <- 90
	allChan <- "name"

	allChan <- Cat{"lucy"}

	// 取出cat结构体，必须要类型断言，否则编译不通过
	<-allChan // 排除前两个
	<-allChan

	newCat, t := <-allChan // 取出的是接口

	fmt.Println("ttt", t)

	cat := newCat.(Cat) // 类型断言
	fmt.Println(cat, " name ：", cat.name)

	allChan <- 1000
	close(allChan) // 必须关闭才能for遍历
	//for e := range allChan {
	//	fmt.Println(e)
	//}

	s, ok := <-allChan
	fmt.Println(s, ok)
	s1, ok1 := <-allChan
	fmt.Println(s1, ok1)
}
