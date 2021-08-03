package main

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
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 1. 尝试连接RabbitMQ，建立连接
	// 该连接抽象了套接字连接，并为我们处理协议版本协商和认证等。
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		failOnError(err, "Failed to connect to RabbitMQ")
	}
	defer conn.Close()

	fmt.Println(conn)
}
