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
	// 这里需要注意，队列申明之后是不能对队列的属性进行修改的，会报错，可以删了重建
	q, err := ch.QueueDeclare(
		"hello", // name
		true,    // durable  是否持久化，重启之后是否保存
		false,   // delete when unused 是否在消费完成后自动删除队列
		false,   // exclusive   //相当于其他连接不能用了： 是否独占队列只对首次声明它的连接（Connection）可见（后面创建的相同名称的队列会报错）， 会在其连接断开的时候自动删除。
		false,   // no-wait  是否非阻塞，true表示是。阻塞：表示创建交换器的请求发送后，阻塞等待RMQ Server返回信息。非阻塞：不会阻塞等待RMQ Server的返回信息，而RMQ Server也不会返回信息。（不推荐使用）
		nil,     // arguments
	)
	if err != nil {
		//fmt.Println("======", err.Error())
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
	//select {}  // 打开可以a看到生产者的conntions和channel
}

func main() {
	for i := 0; i < 10; i++ {
		go Producer(MpUrl)
	}
	select {}
}
