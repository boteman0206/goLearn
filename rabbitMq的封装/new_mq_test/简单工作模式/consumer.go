package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"time"
)

func Consumer(url string) {

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

	// 第三步 申明队列
	q, err := ch.QueueDeclare(
		"hello", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return
	}

	// 第四步： 消费消息
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack 自动确认
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	for d := range msgs {
		strNow := time.Now().Format(time.RFC3339Nano)
		fmt.Printf("消费之收到了消息： Received a message: %s - 时间 %s \n", d.Body, strNow)
	}
}

func main() {

	go Consumer("amqp://noah:noah@localhost:5672/noah_vir")

	select {}
}
