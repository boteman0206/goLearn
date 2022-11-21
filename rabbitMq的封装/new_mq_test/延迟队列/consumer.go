package main

import (
	"github.com/streadway/amqp"
	"log"
)

//=====================================公共参数方法开始===============================================
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
	Queue1      = "Queue1"
	Exchange1   = "Exchange1"
	RoutingKey1 = "RoutingKey1"
)

//=====================================公共参数结束===============================================

func main() {
	// # ========== 1.创建连接 ==========
	mq := NewRabbitMQ("amqp://noah:noah@localhost:5672/noah_vir")
	defer mq.Close()
	mqCh := mq.Channel

	// # ========== 2.消费消息 ==========
	msgsCh, err := mqCh.Consume(Queue1, "", false, false, false, false, nil)
	FailOnError(err, "消费队列失败")

	forever := make(chan bool)
	go func() {
		for d := range msgsCh {
			// 要实现的逻辑
			log.Printf("接收的消息: %s", d.Body)

			// 手动应答
			d.Ack(false)
			//d.Reject(true)
		}
	}()
	log.Printf("[*] Waiting for message, To exit press CTRL+C")
	<-forever
}
