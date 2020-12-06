package main

import (
	"fmt"
	"reflect"
)

// 全局变量加锁不完美，等待时间不好计算
//代码改进在demo3全局变量加锁的基础上使用channel来解决

/**
为何要是用channel，
1：本质是一个数据结构，队列，
2：先进先出
3：线程安全不需要加锁
4：管道是有类型的，string的管道只能放string
5: todo 引用类型
6： 必须初始化才能写入数据，及make后使用

*/
func main() {

	var stringChan chan string
	stringChan = make(chan string, 10)

	fmt.Println("类型： ", reflect.TypeOf(stringChan))
	fmt.Println(stringChan) // 0xc0000180c0地址

	// 向管道写入数据
	stringChan <- "jack"
	stringChan <- "lucy"

	// 查看管道的属性  chan : 2 cap 10
	fmt.Println("chan :", len(stringChan), "cap", cap(stringChan))

	// todo 数据不能超过容量前面申明的10个 取出一个就能在加一个

	// 管道读取数据
	name1 := <-stringChan
	fmt.Println("name1 :", name1)

	// 管道为空，继续取出数据会报错

}
