package main

import (
	"fmt"
	"godemo/src/每日一库练习/rabbitMq的封装/mq集合"
)

func main() {

	// 简单的数据消费
	mq := mq集合.NewRabbitMQSimple("kuteng")

	for i := 0; i < 10; i++ {
		sprintf := fmt.Sprintf("kuteng : %d", i)
		mq.PublishSimple(sprintf)
	}

}
