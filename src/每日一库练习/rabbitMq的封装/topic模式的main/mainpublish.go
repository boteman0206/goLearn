package main

import (
	"fmt"
	"godemo/src/每日一库练习/rabbitMq的封装/topic模式mq"
	"strconv"
	"time"
)

func main() {
	kutengOne := topic模式mq.NewRabbitMQTopic("exKutengTopic", "kuteng.topic.one")
	kutengTwo := topic模式mq.NewRabbitMQTopic("exKutengTopic", "kuteng.topic.two")
	for i := 0; i <= 100; i++ {
		kutengOne.PublishTopic("Hello kuteng topic one!" + strconv.Itoa(i))
		kutengTwo.PublishTopic("Hello kuteng topic Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}

}
