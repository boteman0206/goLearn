package routing模式的mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

const MQURL = "amqp://test01:test01@127.0.0.1:5672/kuteng"

//rabbitMQ结构体
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机名称
	Exchange string
	//bind Key 名称
	Key string
	//连接信息
	Mqurl string
}

//创建结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
}

//断开channel 和 connection
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

//创建RabbitMQ实例
func NewRabbitMQRouting(exchangeName string, routingKey string) *RabbitMQ {
	//创建RabbitMQ实例
	rabbitmq := NewRabbitMQ("", exchangeName, routingKey)
	var err error
	//获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "failed to connect rabbitmq!")
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "failed to open a channel")
	return rabbitmq
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

func (r *RabbitMQ) PublishRouting(message string) {
	// 1 尝试创建交换机
	err := r.channel.ExchangeDeclare(r.Exchange, "direct", true, false, false, false, nil)
	r.failOnErr(err, "declare error")

	// 2：发送消息
	r.channel.Publish(r.Exchange, r.Key, false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(message)})

}

//路由模式接受消息

func (r *RabbitMQ) RecieveRouting() {
	//1： 试探性创建链接
	err := r.channel.ExchangeDeclare(r.Exchange, "direct", true, false, false, false, nil)

	r.failOnErr(err, "recieve err")
	// 2 试探性的创建队列 这里注意队列的名称不要写
	q, err := r.channel.QueueDeclare(
		"", //随机生产队列名称
		false,
		false,
		true,
		false,
		nil,
	)

	r.failOnErr(err, "queue error")

	// 3 绑定队列到exchange
	err = r.channel.QueueBind(q.Name, r.Key, r.Exchange, false, nil)

	// 消费消息
	messages, err := r.channel.Consume(q.Name, "", true, false, false, false, nil)

	bools := make(chan bool)
	go func() {
		for e := range messages {
			log.Printf("Received a message : %s", e.Body)

		}
	}()

	<-bools

}
