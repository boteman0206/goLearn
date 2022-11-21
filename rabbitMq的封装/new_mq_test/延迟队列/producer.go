package main

/**
延迟队列
	延迟队列：队列中的元素需要在指定时间取出和处理。例如，用户发起订单，十分钟内未支付则自动取消。
	当数据量很大时，采取轮询的方式显然是不合理的，会给数据库带来很大压力


RabbitMQ中的TTL
	TTL，最大存活时间，表明消息或该队列中所有消息的最大存活时间。
有两种方式设置：
	1: 针对每条信息设置TTL
		消息即使过期，也不一定会被马上丢弃，因为消息是否过期是在即将投递到消费者之前判定的。
		即RabbitMQ只对处于队头的消息判断是否过期（即不会扫描队列），所以，很可能队列中已存在死消息，但是队列并不知情。这会影响队列统计数据的正确性，妨碍队列及时释放资源。
	2: 在创建队列时设置队列的x-message-ttl属性
		如果设置了队列的 TTL 属性，那么一旦消息过期，就会被队列丢弃(如果配置了死信队列被丢到死信队列中)。  设置队列的过期时间是对该队列的所有消息生效的。

实现延迟队列的方式：
	一：延时队列核心 = 死信队列 + TTL：TTL让消息延迟多久后成为死信，消费者一直处理死信队列里的信息就行。（可以参考死信队列的实现）
	二：通过插件实现延迟队列
		此插件的原理是将消息在交换机处暂存储在mnesia(一个分布式数据系统)表中，延迟投递到队列中，等到消息到期再投递到队列当中。


本文主要是通过插件来实现：
	https://blog.csdn.net/qq_36551991/article/details/107213281 插件安装说明博客



*/

//这里只针对插件实现
import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"time"
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

/**
插件生产者实现的关键点：
	1.在声明交换机时不在是direct类型，而是x-delayed-message类型，这是由插件提供的类型；
	2.交换机要增加"x-delayed-type": "direct"参数设置；
	3.发布消息时，要在 Headers 中设置x-delay参数，来控制消息从交换机过期时间；
*/
func producer() {
	// # ========== 1.创建连接 ==========
	mq := NewRabbitMQ("amqp://noah:noah@localhost:5672/noah_vir")
	defer mq.Close()
	mqCh := mq.Channel

	// # ========== 2.设置队列（队列、交换机、绑定） ==========
	// 声明队列
	var err error
	_, err = mqCh.QueueDeclare(Queue1, true, false, false, false, amqp.Table{
		// "x-message-ttl": 60000, // 消息过期时间（队列级别）,毫秒
	})
	FailOnError(err, "创建队列失败")

	// 声明交换机
	//err = mqCh.ExchangeDeclare(Exchange1, amqp.ExchangeDirect, true, false, false, false, nil) todo 在声明交换机时不在是direct类型，而是x-delayed-message类型，这是由插件提供的类型；
	err = mqCh.ExchangeDeclare(Exchange1, "x-delayed-message", true, false, false, false, amqp.Table{
		"x-delayed-type": "direct", // todo  交换机要增加"x-delayed-type": "direct"参数设置；
	})
	FailOnError(err, "创建交换机失败")

	// 队列绑定（将队列、routing-key、交换机三者绑定到一起）
	err = mqCh.QueueBind(Queue1, RoutingKey1, Exchange1, false, nil)
	FailOnError(err, "队列、交换机、routing-key 绑定失败")

	// # ========== 4.发布消息 ==========
	message := "msg" + strconv.Itoa(int(time.Now().Unix()))
	fmt.Println(message)
	// 发布消息
	err = mqCh.Publish(Exchange1, RoutingKey1, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
		//Expiration: "10000", // 消息过期时间（消息级别）,毫秒
		Headers: map[string]interface{}{
			"x-delay": "5000", // todo 消息从交换机过期时间,毫秒（x-dead-message插件提供）
		},
	})
	FailOnError(err, "消息发布失败")
}

func main() {
	go producer()
	select {}
}
