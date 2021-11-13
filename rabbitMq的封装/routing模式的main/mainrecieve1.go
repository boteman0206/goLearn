package main

import (
	"godemo/rabbitMq的封装/routing模式的mq"
)

func main() {
	kutengone := routing模式的mq.NewRabbitMQRouting("kuteng", "kuteng_one")
	kutengone.RecieveRouting()
}
