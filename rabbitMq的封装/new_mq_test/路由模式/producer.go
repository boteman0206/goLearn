package main

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/streadway/amqp"
	"log"
	"os"
)

/*
http://rabbitmq.mr-ping.com/tutorials_with_golang/[4]Routing.html


RabbitMq消息模式的核心思想是：一个生产者并不会直接往一个队列中发送消息，事实上，生产者根本不知道它发送的消息将被转发到哪些队列。
实际上，生产者只能把消息发送给一个exchange，exchange只做一件简单的事情：一方面它们接收从生产者发送过来的消息，另一方面，它们把接收到的消息推送给队列。一个exchage必须清楚地知道如何处理一条消息。


// todo 1.声明队列, 一下很多例子都是乙方申明队列的
	/*
		如果只有一方声明队列，可能会导致下面的情况：
			a)消费者是无法订阅或者获取不存在的MessageQueue中信息
			b)消息被Exchange接受以后，如果没有匹配的Queue，则会被丢弃

		为了避免上面的问题，所以最好选择两方一起声明
		ps:如果客户端尝试建立一个已经存在的消息队列，Rabbit MQ不会做任何事情，并返回客户端建立成功的
*/




*/

/**
读取参数
第一个参数为文件名称，第二个参数为输入的信息 [tets.exe abc]
*/
func severityFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "info"
	} else {
		s = os.Args[1]
	}
	return s
}

func Producer(url string, num int) {

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

	err = ch.ExchangeDeclare(
		"logs_direct", // name 交换机名称
		"direct",      // type 直连交换机
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)

	// 基于交换机的不需要申明队列，直接将消息发到交换机上面

	body := "t am direct work  " + cast.ToString(num)
	err = ch.Publish(
		"logs_direct",         // exchange
		severityFrom(os.Args), // routing key 更具输入的参数来发送消息到队列 默认是info队列
		false,                 // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, // 消息的持久性
			ContentType:  "text/plain",
			Body:         []byte(body),
		})

	if err != nil {
		fmt.Println("err: ", err.Error())
		return
	}
	log.Printf(" [x] Sent %s\n", body)
}

func main() {
	for i := 0; i < 3; i++ {
		Producer("amqp://noah:noah@localhost:5672/noah_vir", i)
	}

	select {}
}
