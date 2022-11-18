package main

import (
	"github.com/streadway/amqp"
	"log"
)

/*
	简单工作模式和work工作模式发送方本质是相同的，不使用交换机直接面向队列的。实际上用的是默认的：路由键和队列名称相同

*/

var (
	MpUrl = "amqp://noah:noah@localhost:5672/noah_vir"
)

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

	//第三步： declare a queue 如果是其他的工作模式还需要申明交互机，这里只是简单工作模式
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return
	}

	// 第四步  发送消息到消息队列
	body := "Hello noah!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return
	}

	log.Printf(" [x] Sent %s\n", body)

}

func main() {
	for i := 0; i < 10; i++ {
		Producer(MpUrl)
	}
}
