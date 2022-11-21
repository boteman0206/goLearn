package main

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/streadway/amqp"
	"log"
	"strconv"
)

/**
RabbitMQ3.5.0之后官方版本已经实现了优先级队列。数值越大则优先级越高。

创建优先级队列，需要增加x-max-priority参数，指定一个优先级的数值大小，这里最好是0～10之间，用来表示这个queue的最大优先级。（备注：因为生产者和消费者都需要对queue进行声明，所以它们都需要设置这个参数）

生产者在发送消息的时候，需要设置priority属性，最好不要超过上面指定的最大的优先级，一旦超过了这个优先级，发送设置的优先级就不再生效了。在这个范围内的优先级，数字越大，优先级越高。

优先级队列处理的场景，是针对的生产者生产快，消费者消费慢，反之没有意义，毕竟只有queue中有消息堆积的时候，才会需要根据优先级策略进行调度。

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

func producer(num int) {

	// # ========== 1.创建连接 ==========
	mq := NewRabbitMQ("amqp://noah:noah@localhost:5672/noah_vir")
	defer mq.Close()
	mqCh := mq.Channel

	var args amqp.Table
	args = amqp.Table{"x-max-priority": int32(10)}

	//创建队列
	q, err := mqCh.QueueDeclare("priqueue", true, false, false, false, args)
	FailOnError(err, "创建优先级队列失败")

	message := "msg" + strconv.Itoa(num) + "  " + cast.ToString(uint8(num/2+1))
	fmt.Println(message)

	//err = mqCh.Qos(1, 0, false)
	//if err != nil {
	//	fmt.Println("err Qos : ", err.Error())
	//	return
	//}

	// 发布消息
	err = mqCh.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
		Priority:    uint8(num/2 + 1),
	})
	FailOnError(err, "消息发布失败")

}

func main() {
	for i := 0; i < 10; i++ {
		go producer(i)
	}

	select {}
}
