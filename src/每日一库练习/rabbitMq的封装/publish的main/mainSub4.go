package main

import "godemo/src/每日一库练习/rabbitMq的封装/publish工作mq集合"

func main() {
	rabbitmq := publish工作mq集合.NewRabbitMQPubSub("newProduct")
	rabbitmq.RecieveSub()
}
