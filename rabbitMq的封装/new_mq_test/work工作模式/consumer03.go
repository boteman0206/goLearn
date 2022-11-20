package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"time"
)

/**
公平调度的消费者，设置预取值
	 err = ch.Qos(
                1,     // prefetch count
                0,     // prefetch size
                false, // global
        )

prefetchSize：可接收消息的大小的
prefetchCount：处理消息最大的数量。举个例子，如果输入1，那如果接收一个消息，但是没有应答，
	则客户端不会收到下一个消息，消息只会在队列中阻塞。如果输入3，那么可以最多有3个消息不应答，
	如果到达了3个，则发送端发给这个接收方得消息只会在队列中，而接收方不会有接收到消息的事件产生。
	总结说，就是在下一次发送应答消息前，客户端可以收到的消息最大数量。
global：是不是针对整个Connection的，因为一个Connection可以有多个Channel，
	  如果是false则说明只是针对于这个Channel的


使用消息响应和prefetch_count你就可以搭建起一个工作队列了。这些持久化的选项使得在RabbitMQ重启之后仍然能够恢复。


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

	// 第三步 申明队列
	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return
	}

	//设置预取值,一次只往队列里里面发送一条消息，能者多劳
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	// 第四步： 消费消息
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack 需要手动确认你才会删除消息
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	for d := range msgs {
		time.Sleep(1 * time.Second)
		strNow := time.Now().Format(time.RFC3339Nano)
		fmt.Printf("消费之收到了消息： Received a message: %s - 时间 %s \n", d.Body, strNow)

		d.Ack(false) // 手动确认消息

	}
}

func main() {

	go Consumer("amqp://noah:noah@localhost:5672/noah_vir")

	select {}
}
