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

topic交换机
	1: 发送到topic交换机的消息不可以携带随意routing_key，它的routing_key必须是一个由.分隔开的词语列表。
	2: binding key也必须拥有同样的格式。topic交换机背后的逻辑跟direct交换机很相似 —— 一个携带着特定routing_key的消息会被topic交换机投递给绑定键与之想匹配的队列。但是它的binding key和routing_key有两个特殊应用方式：
		* (星号) 用来表示一个单词.
		# (井号) 用来表示任意数量（零个或多个）单词。

和直连交换机的联系
	Topic交换机是很强大的，它可以表现出跟其他交换机类似的行为 当一个队列的binding key为 "#"（井号） 的时候，这个队列将会无视消息的routing key，接收所有的消息。
	当 * (星号) 和 # (井号) 这两个特殊字符都未在binding key中出现的时候，此时Topic交换机就拥有的direct交换机的行为。



绑定键为 * 的队列会取到一个routing key为空的消息吗？
	测试过，可以
a.*.# 和 a.#的区别在哪儿？
   a.*.# 必须是两个单词组成的路由才可以，而a.#一个单词即可
绑定键为 #.* 的队列会获取到一个名为..的路由键的消息吗？它会取到一个routing key为单个单词的消息吗？
	"#.*"测试过两个都是可以的


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
		"logs_topic", // name 交换机名称
		"topic",      // type 直连交换机
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)

	// 基于交换机的不需要申明队列，直接将消息发到交换机上面

	body := "t am topic work  " + cast.ToString(num)
	err = ch.Publish(
		"logs_topic",          // exchange
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
	log.Printf(" [x] Sent topic %s\n", body)
}

func main() {
	for i := 0; i < 3; i++ {
		Producer("amqp://noah:noah@localhost:5672/noah_vir", i)
	}

	select {}
}

/**
启动方式：
	生产者：
		go run .\consumer.go *.*.rabbit
	 	go run .\consumer.go *.red.pig

	消费者：
		go run .\consumer.go *.*.rabbit

*/
