package main

import "github.com/streadway/amqp"

/**
工作队列的核心思想：
	对于生产者（产生消息的人）：避免必须立刻执行“资源紧张”的任务。
	对于消息队列：生产者想要做的“任务”会被封装成一个消息放在队列里。
	对于消费者（处理任务的人）：当你有多个“工人”时，这些任务会被轮询分配给不同的工人。
	这个思想也被应用于那些需要处理不能在一个很短的HTTP请求窗口期间完成的复杂任务的网页程序中。

*/

func Producer(url string) {

	// 第一步骤  连接RabbitMQ服务器
	conn, err := amqp.Dial(url)
	if err != nil {
		return
	}
	defer conn.Close()

	//第二步： 连接信道channel 我们创建一个通道，这是大部分用于完成工作的API所在的地方
	ch, err := conn.Channel()
	if err != nil {
		return
	}
	defer ch.Close()

}

func main() {
	for i := 0; i < 5; i++ {
		Producer("amqp://noah:noah@localhost:5672/noah_vir")
	}
}
