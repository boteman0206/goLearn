package main

import (
	"fmt"
	"godemo/rabbitMq的封装/simple和work工作模式mq集合"
)

func main() {

	// 简单的数据消费
	mq := simple和work工作模式mq集合.NewRabbitMQSimple("kuteng")

	for i := 0; i < 100; i++ {
		sprintf := fmt.Sprintf("kuteng : %d", i)
		mq.PublishSimple(sprintf)
	}

}
