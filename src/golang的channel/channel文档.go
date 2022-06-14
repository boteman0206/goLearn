package main

import (
	"fmt"
	"time"
)

/**
https://juejin.cn/post/6844904016254599176 参考
https://github.com/golang/go/blob/master/src/runtime/chan.go 源码
1： hchan的结构
type hchan struct {
   qcount   uint           //  当前队列中剩余元素个数
   dataqsiz uint           // 环形队列长度，即缓冲区的大小，即make（chan T，N），N.
   buf      unsafe.Pointer // 环形队列指针
   elemsize uint16 //  每个元素的大小
   closed   uint32  // 表示当前通道是否处于关闭状态。创建通道后，该字段设置为0，即通道打开; 通过调用close将其设置为1，通道关闭。
   elemtype *_type //  元素类型，用于数据传递过程中的赋值；
   sendx    uint   // sendx uint和recvx uint是环形缓冲区的状态字段，它指示缓冲区的当前索引 - 支持数组，它可以从中发送数据和接收数据。
   recvx    uint   // 和上面的一样
   recvq    waitq  // list of recv waiters 等待读消息的goroutine队列
   sendq    waitq  // 等待写消息的goroutine队列
   lock mutex 互斥锁，为每个读写操作锁定通道，因为发送和接收必须是互斥操作。
}
type waitq struct {
    first *sudog
    last  *sudog
}


2：创建channel 有两种，一种是带缓冲的channel，一种是不带缓冲的channel
// 带缓冲
ch := make(chan Task, 3)
// 不带缓冲
ch := make(chan int)


3： 什么情况下关闭 channel 会造成 panic ？
	// 1.未初始化时关闭
	// 2.重复关闭
	// 3.关闭后发送
	// 4.发送时关闭

4：关闭后的通道有以下特点：
 	1.对一个关闭的通道再发送值就会导致panic。
    2.对一个关闭的通道进行接收会一直获取值直到通道为空。
    3.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
    4.关闭一个已经关闭的通道会导致panic。


5: 无缓冲通道
func main() {
    ch := make(chan int)
    ch <- 10
    fmt.Println("发送成功")
}
todo fatal error: all goroutines are asleep - deadlock!
因为我们使用ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有人接收值的时候才能发送值
就像你住的小区没有快递柜和代收点，快递员给你打电话必须要把这个物品送到你的手中，简单来说就是无缓冲的通道必须有接收才能发送
上面的代码会阻塞在ch <- 10这一行代码形成死锁，那如何解决这个问题呢？
无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时值才能发送成功，两个goroutine将继续执行。相反，如果接收操作先执行，接收方的goroutine将阻塞，直到另一个goroutine在该通道上发送一个值。
使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。因此，无缓冲通道也被称为同步通道


6: 有缓冲的通道
解决上面问题的方法还有一种就是使用有缓冲区的通道。
只要通道的容量大于零，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素的数量。就像你小区的快递柜只有那么个多格子，格子满了就装不下了，就阻塞了，等到别人取走一个快递员就能往里面放一个。

7: 关闭通道 close()

8：如何优雅的从通道循环取值
当通过通道发送有限的数据时，我们可以通过close函数关闭通道来告知从该通道接收值的goroutine停止等待。当通道被关闭时，往该通道发送值会引发panic，从该通道里接收的值一直都是类型零值。那如何判断一个通道是否被关闭了呢？


9: for-range遍历
在遍历时，如果channel 没有关闭，那么会一直等待下去，出现 deadlock 的错误；如果在遍历时channel已经关闭，那么在遍历完数据后自动退出遍历。也就是说，for range 的遍历方式时阻塞型的遍历方式。


10： for select select可以处理非阻塞式消息发送、接收及多路选择。
select中有case代码块，用于channel发送或接收消息，任意一个case代码块准备好时，执行其对应内容；
多个case代码块准备好时，随机选择一个case代码块并执行；所有case代码块都没有准备好，则等待；
还可以有一个default代码块，所有case代码块都没有准备好时执行default代码块。





*/

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}
func main() {
	ch := make(chan int)
	go recv(ch) // 启用goroutine从通道接收值
	//ch <- 10
	fmt.Println("发送成功")
	time.Sleep(40 * time.Second)
}
