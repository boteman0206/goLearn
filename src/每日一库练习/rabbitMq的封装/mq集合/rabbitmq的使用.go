package mq集合

/**
http://www.topgoer.com/%E6%95%B0%E6%8D%AE%E5%BA%93%E6%93%8D%E4%BD%9C/go%E6%93%8D%E4%BD%9CRabbitMQ/RabbitMQ%E5%AE%89%E8%A3%85.html

todo 如果rabbitmq启动正常但是无法通过链接访问  http://localhost:15672/
	1： 打开RabbitMQ节点： rabbitmqctl start_app
	2： 开启RabbitMQ管理模块的插件，并配置到RabbitMQ节点上 rabbitmq-plugins enable rabbitmq_management
	3： 关闭rabbitMQ节点 rabbitmqctl stop
	4: 重新启动     //参考解决 https://blog.csdn.net/sdgames/article/details/104267859

*/
import (
	"fmt"
	"github.com/streadway/amqp"
)

//直接模式队列生产
func (r *RabbitMQ) PublishSimple(message string) {
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	_, err := r.channel.QueueDeclare(
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
	//调用channel 发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		//如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}
