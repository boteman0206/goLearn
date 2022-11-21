package main

import (
	"github.com/streadway/amqp"
	"log"
)

/**
由于队列、交换机都交由生产者来创建，消费者只需做两件，一是建立连接、二是消费消息。
也由于这个原因，消费者要晚于生产者启动，可以保证消费的时候，队列是存在的。

*/
type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func (c *RabbitMQ) Close() {
	c.Conn.Close()
	c.Channel.Close()
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func NewRabbitMQ(MqUrl string) *RabbitMQ {
	conn, err := amqp.Dial(MqUrl)
	FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")

	return &RabbitMQ{
		Conn:    conn,
		Channel: ch,
	}
}

var (
	NormalQueue      = "normal_queue"       //  正常队列
	NormalExchange   = "normal_exchange"    // 正常交换机
	NormalRoutingKey = "normal_routing_key" // 正常路由key

	DeadExchange   = "dead_exchange"    // 死信交换机
	DeadRoutingKey = "dead_routing_key" // 死信路由key
	DeadQueue      = "dead_queue"       // 死信队列

)

//正常 消费者
func main_normal() {
	// # ========== 1.创建连接 ==========
	mq := NewRabbitMQ("amqp://noah:noah@localhost:5672/noah_vir")
	defer mq.Close()
	mqCh := mq.Channel

	// # ========== 2.消费消息 ==========
	msgsCh, err := mqCh.Consume(NormalQueue, "", false, false, false, false, nil)
	FailOnError(err, "消费normal队列失败")

	forever := make(chan bool)
	go func() {
		for d := range msgsCh {
			// 要实现的逻辑
			log.Printf("normal 接收的消息: %s", d.Body)

			// 手动应答
			d.Ack(false)
			//d.Reject(true)
		}
	}()
	log.Printf("[*] Waiting for message, To exit press CTRL+C")
	<-forever
}

func main_dead() {
	// # ========== 1.创建连接 ==========
	mq := NewRabbitMQ("amqp://noah:noah@localhost:5672/noah_vir")
	defer mq.Close()
	mqCh := mq.Channel

	// # ========== 2.消费死信消息 ==========
	msgsCh, err := mqCh.Consume(DeadQueue, "", false, false, false, false, nil)
	FailOnError(err, "消费dead队列失败")

	forever := make(chan bool)
	go func() {
		for d := range msgsCh {
			// 要实现的逻辑
			log.Printf("dead 接收的消息: %s", d.Body)

			// 手动应答
			d.Ack(false)
			//d.Reject(true)
		}
	}()
	log.Printf("[*] Waiting for message, To exit press CTRL+C")
	<-forever
}

func main() {
	//go main_normal()
	go main_dead()
	select {}
}
