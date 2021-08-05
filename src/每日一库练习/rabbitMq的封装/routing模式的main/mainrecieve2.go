package main

import (
	"godemo/src/每日一库练习/rabbitMq的封装/routing模式的mq"
)

func main() {
	kutengtwo := routing模式的mq.NewRabbitMQRouting("kuteng", "kuteng_two")
	kutengtwo.RecieveRouting()
}
