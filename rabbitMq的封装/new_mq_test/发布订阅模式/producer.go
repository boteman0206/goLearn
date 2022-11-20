package main

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/streadway/amqp"
	"log"
)

/*
http://rabbitmq.mr-ping.com/tutorials_with_golang/[3]Publish_Subscribe.html

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
