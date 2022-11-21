package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"time"
)

/**
死信：由于某些原因（消息TTL过期、队列达到最大长度、消息被拒绝）导致队列中的消息无法被处理。
	RabbitMQ死信队列机制：当消息消费发生异常时，将消息投入死信队列。（例如，用户下单成功但未在指定时间内支付 -> 消息自动失效）
	（1）消息被否定确认，使用 channel.basicNack 或 channel.basicReject ，并且此时requeue 属性被设置为false。
	（2）消息在队列的存活时间超过设置的TTL时间。
	 (3）消息队列的消息数量已经超过最大队列长度。

x-message-ttl             5000,           // 消息过期时间,毫秒
x-max-length               6              指定队列的长度


死信消息变化
	如果队列配置了参数 x-dead-letter-routing-key 的话，“死信”的路由key将会被替换成该参数对应的值。如果没有设置，则保留该消息原有的路由key
	由于被抛到了死信交换机，所以消息的Exchange Name也会被替换为死信交换机的名称。
log.info("死信消息properties：{}", message.getMessageProperties()); java代码获取死信消息属性
	x-first-death-exchange  第一次被抛入的死信交换机的名称
	x-first-death-reason   第一次成为死信的原因，rejected：消息在重新进入队列时被队列拒绝，由于default-requeue-rejected 参数被设置为false。expired ：消息过期。maxlen ： 队列内消息数量超过队列最大容量
	x-first-death-queue    第一次成为死信前所在队列名称
	x-death    历次被投入死信交换机的信息列表，同一个消息每次进入一个死信交换机，这个数组的信息就会被更新

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

func main() {

	// # ========== 1.创建连接 ==========
	mq := NewRabbitMQ("amqp://noah:noah@localhost:5672/noah_vir")
	defer mq.Close()
	mqCh := mq.Channel

	// # ========== 2.设置队列（队列、交换机、绑定） ==========
	var err error
	_, err = mqCh.QueueDeclare(NormalQueue, true, false, false, false, amqp.Table{
		"x-message-ttl":             5000,           // 消息过期时间,毫秒
		"x-dead-letter-exchange":    DeadExchange,   // 指定死信交换机
		"x-dead-letter-routing-key": DeadRoutingKey, // 指定死信routing-key
	})

	FailOnError(err, "创建normal队列失败")

	//目前，普通队列和交换机都已经创建，但它们都是独立存在，没有关联。
	err = mqCh.ExchangeDeclare(NormalExchange, amqp.ExchangeDirect, true, false, false, false, nil)
	FailOnError(err, "创建normal交换机失败")

	//通过 QueueBind 将队列、routing-key、交换机三者绑定到一起。
	err = mqCh.QueueBind(NormalQueue, NormalRoutingKey, NormalExchange, false, nil)
	FailOnError(err, "normal：队列、交换机、routing-key 绑定失败")

	// # ========== 3.设置死信队列（队列、交换机、绑定） ==========
	/**
	设置死信队列（队列、交换机、绑定）
	同样死信队列，也需要创建队列、创建交换机和绑定。
	当死信队列建立完毕，普通队列通过 x-dead-letter-exchange 和 x-dead-letter-routing-key 参数的指定，便可生效，死信队列便与普通队列连通。
	*/
	// 声明死信队列
	// args 为 nil。切记不要给死信队列设置消息过期时间,否则失效的消息进入死信队列后会再次过期。
	_, err = mqCh.QueueDeclare(DeadQueue, true, false, false, false, nil)
	FailOnError(err, "创建dead队列失败")

	// 声明交换机
	err = mqCh.ExchangeDeclare(DeadExchange, amqp.ExchangeDirect, true, false, false, false, nil)
	FailOnError(err, "创建dead队列失败")

	// 队列绑定（将队列、routing-key、交换机三者绑定到一起）
	err = mqCh.QueueBind(DeadQueue, DeadRoutingKey, DeadExchange, false, nil)
	FailOnError(err, "dead：队列、交换机、routing-key 绑定失败")

	message := "msg" + strconv.Itoa(int(time.Now().Unix()))
	fmt.Println(message)

	// # ========== 4.发布消息 ==========
	// 发布消息
	err = mqCh.Publish(NormalExchange, NormalRoutingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})
	FailOnError(err, "消息发布失败")

}
