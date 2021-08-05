package main

import (
	"godemo/src/每日一库练习/rabbitMq的封装/routing模式的mq"
)

func main() {
	kutengone := routing模式的mq.NewRabbitMQRouting("kuteng", "kuteng_one")
	kutengone.RecieveRouting()
}
