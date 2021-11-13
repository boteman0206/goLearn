package main

import (
	"godemo/rabbitMq的封装/topic模式mq"
)

func main() {
	kutengOne := topic模式mq.NewRabbitMQTopic("exKutengTopic", "#")
	kutengOne.RecieveTopic()
}
