package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"time"
)

/**
普通消费者
*/
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

	var args amqp.Table
	args = amqp.Table{"x-max-priority": int32(10)}
	q, err := ch.QueueDeclare("priqueue", true, false, false, false, args)
	if err != nil {
		fmt.Println("err QueueDeclare : ", err.Error())
		return
	}

	// 一次一次的消费
	err = ch.Qos(1, 0, false)
	if err != nil {
		fmt.Println("err Qos : ", err.Error())
		return
	}

	// 第四步： 消费消息
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack 自动确认
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	for d := range msgs {
		time.Sleep(3 * time.Second)
		strNow := time.Now().Format(time.RFC3339Nano)
		fmt.Printf("优先级队列： Received a message: %s - 时间 %s \n", d.Body, strNow)
		err := d.Ack(false)
		if err != nil {
			fmt.Println("ack err : ", err.Error())
			return
		}
	}
}

func main() {

	go Consumer("amqp://noah:noah@localhost:5672/noah_vir")

	select {}
}
