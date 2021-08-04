package mq集合

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

//连接信息amqp://test01:test01@localhost:5672/kuteng这个信息是固定不变的amqp://是固定参数，后面两个是用户名密码ip地址端口号Virtual Host
const MQURL = "amqp://test01:test01@localhost:5672/kuteng"

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

// 创建简单模式的导入
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	// 创建rabbitmq的实例
	mq := NewRabbitMQ(queueName, "", "")
	var err error
	// 获取connection链接
	mq.conn, err = amqp.Dial(mq.Mqurl)
	mq.failOnErr(err, "failed to connect rabbit")

	// 获取channel
	mq.channel, err = mq.conn.Channel()
	mq.failOnErr(err, "failed to open channel")
	return mq
}

//simple 模式下消费者
func (r *RabbitMQ) ConsumeSimple(consumerName string) {
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	//接收消息
	msgs, err := r.channel.Consume(
		q.Name, // queue
		//用来区分多个消费者
		"", // consumer
		//是否自动应答
		true, // auto-ack
		//是否独有
		false, // exclusive
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-local
		//列是否阻塞
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	//启用协程处理消息
	go func() {
		for d := range msgs {
			//消息逻辑处理，可以自行设计逻辑
			log.Printf("Received a message: %s  %s", d.Body, consumerName)

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
