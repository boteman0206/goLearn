package main

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/streadway/amqp"
	"log"
)

/**
工作队列的核心思想：
	对于生产者（产生消息的人）：避免必须立刻执行“资源紧张”的任务。
	对于消息队列：生产者想要做的“任务”会被封装成一个消息放在队列里。
	对于消费者（处理任务的人）：当你有多个“工人”时，这些任务会被轮询分配给不同的工人。
	这个思想也被应用于那些需要处理不能在一个很短的HTTP请求窗口期间完成的复杂任务的网页程序中。


todo
注意点： work工作模式的代码其实和简单工作模式的代码其实是一样的，只不过一个需要启动多个消费者来消费而已
consumer01为普通普通消费者
consumer02为需要手动确认ack的消费者
consumer03为设置预取值选项


消息确认
	为了防止消息丢失，RabbitMQ提供了消息响应（acknowledgments）。消费者会通过一个ack（响应），告诉RabbitMQ已经收到并处理了某条消息，然后RabbitMQ就会释放并删除这条消息。
	如果消费者（consumer）挂掉了，没有发送响应，RabbitMQ就会认为消息没有被完全处理，然后重新发送给其他消费者（consumer）。这样，即使工作者（workers）偶尔的挂掉，也不会丢失消息。


忘记确认
	忘记ack是一个常见的错误。这是一个简单的错误，但后果是严重的。当客户端退出时，消息将被重新传递（这可能看起来像随机重新传递），
	但是RabbitMQ将会占用越来越多的内存，因为它无法释放任何未经消息的消息 为了排除这种错误，你可以使用rabbitmqctl命令，输出messages_unacknowledged字段：
		linux：	  sudo rabbitmqctl list_queues name messages_ready messages_unacknowledged
		Windows上执行：rabbitmqctl.bat list_queues name messages_ready messages_unacknowledged

如何保证消息不丢失
消息持久化： 如果你没有特意告诉RabbitMQ，那么在它退出或者崩溃的时候，将会丢失所有队列和消息。为了确保信息不会丢失，有两个事情是需要注意的：我们必须把“队列”和“消息”设为持久化。
1：为了不让队列消失，需要把队列声明为持久化（durable）：
2： 需要将消息标记为持久性 - 通过设置amqp.Publishing的amqp.Persistent属性完成。
	现在我们需要将消息标记为持久性 - 通过设置amqp.Publishing的amqp.Persistent属性完成。

todo 注意：消息持久化
将消息设为持久化并不能完全保证不会丢失。以上代码只是告诉了RabbitMq要把消息存到硬盘，
但从RabbitMq收到消息到保存之间还是有一个很小的间隔时间。因为RabbitMq并不是所有的消息都使用fsync(2)——
它有可能只是保存到缓存中，并不一定会写到硬盘中。并不能保证真正的持久化，但已经足够应付我们的简单工作队列。
如果您需要更强的保证，那么您可以使用publisher confirms.。

公平调度
	我们可以设置预取计数值为1。告诉RabbitMQ一次只向一个worker发送一条消息。换句话说，在处理并确认前一个消息之前，不要向工作人员发送新消息。


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
	body := "t am work " + cast.ToString(num)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
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
