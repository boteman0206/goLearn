package main

import (
	"fmt"
	"godemo/rabbitMq的封装/routing模式的mq"
	"strconv"
	"time"
)

func main() {
	kutengone := routing模式的mq.NewRabbitMQRouting("kuteng", "kuteng_one")
	kutengtwo := routing模式的mq.NewRabbitMQRouting("kuteng", "kuteng_two")
	for i := 0; i <= 100; i++ {
		kutengone.PublishRouting("Hello kuteng one!" + strconv.Itoa(i))
		kutengtwo.PublishRouting("Hello kuteng Two!" + strconv.Itoa(i))
		time.Sleep(2 * time.Second)
		fmt.Println(i)
	}

}
