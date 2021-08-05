package main

import "godemo/src/每日一库练习/rabbitMq的封装/topic模式mq"

func main() {
	kutengOne := topic模式mq.NewRabbitMQTopic("exKutengTopic", "kuteng.*.two")
	kutengOne.RecieveTopic()
}
