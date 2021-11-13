package main

import "godemo/rabbitMq的封装/publish工作mq集合"

func main() {
	rabbitmq := publish工作mq集合.NewRabbitMQPubSub("newProduct")
	rabbitmq.RecieveSub()
}
