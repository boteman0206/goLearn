package main

import (
	"fmt"
	"godemo/rabbitMq的封装/publish工作mq集合"
	"strconv"
	"time"
)

func main() {
	rabbitmq := publish工作mq集合.NewRabbitMQPubSub("newProduct")
	for i := 0; i < 100; i++ {
		rabbitmq.PublishPub("订阅模式生产第" +
			strconv.Itoa(i) + "条" + "数据")
		fmt.Println("订阅模式生产第" +
			strconv.Itoa(i) + "条" + "数据")
		time.Sleep(1 * time.Second)
	}

}
