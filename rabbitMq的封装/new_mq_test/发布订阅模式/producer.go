package main

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/streadway/amqp"
	"log"
)

/*
http://rabbitmq.mr-ping.com/tutorials_with_golang/[3]Publish_Subscribe.html


RabbitMq消息模式的核心思想是：一个生产者并不会直接往一个队列中发送消息，事实上，生产者根本不知道它发送的消息将被转发到哪些队列。
实际上，生产者只能把消息发送给一个exchange，exchange只做一件简单的事情：一方面它们接收从生产者发送过来的消息，另一方面，它们把接收到的消息推送给队列。一个exchage必须清楚地知道如何处理一条消息。


fanout总结
1.生产者不是直接操作队列，而是将数据发送给交换机，由交换机将数据发送给与之绑定的队列。
2：必须声明交换机，并且设置模式：channel.exchangeDeclare(EXCHANGE_NAME, "fanout")，
	其中 fanout 指分发模式（将每一条消息都发送到与交换机绑定的队列）。
3： 队列必须绑定交换机：channel.queueBind(QUEUE_NAME, EXCHANGE_NAME, "");
*/

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
		"fanout_log", // name
		"fanout",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)

	// fanout可以不需要申明队列
	//q, err := ch.QueueDeclare(
	//	"", // name
	//	true,         // durable
	//	false,        // delete when unused
	//	false,        // exclusive
	//	false,        // no-wait
	//	nil,          // arguments
	//)

	body := "t am work " + cast.ToString(num)
	err = ch.Publish(
		"fanout_log", // exchange
		"",           // routing key
		false,        // mandatory
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
	for i := 0; i < 10; i++ {
		Producer("amqp://noah:noah@localhost:5672/noah_vir", i)
	}

	select {}
}
