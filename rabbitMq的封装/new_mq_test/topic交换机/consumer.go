package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
	"time"
)

/**
普通消费者
*/
func severityFrom(args []string) string {
	var s string
	fmt.Println("args: ", args)
	if (len(args) < 2) || os.Args[1] == "" {
		s = "info"
	} else {
		s = os.Args[1]
	}
	return s
}

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

	// 第三步 申明交换机
	err = ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)

	// 第四步 申明队列，这里fanout的使用的是临时队列，生产我们可以自己考虑
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return
	}

	from := severityFrom(os.Args)
	fmt.Println("topic交换机：", from)
	// 绑定队列
	err = ch.QueueBind(
		q.Name,       // queue name
		"#.*",        // routing key
		"logs_topic", // exchange
		false,
		nil,
	)

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
		time.Sleep(3 * time.Second)
		strNow := time.Now().Format(time.RFC3339Nano)
		fmt.Printf("topic模式： Received a message: %s - 时间 %s \n", d.Body, strNow)
	}
}

func main() {

	go Consumer("amqp://noah:noah@localhost:5672/noah_vir")

	select {}
}
