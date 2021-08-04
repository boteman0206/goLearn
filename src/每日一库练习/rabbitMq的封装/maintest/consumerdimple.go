package main

import (
	"fmt"
	"godemo/src/每日一库练习/rabbitMq的封装/mq集合"
	"time"
)

func consumer1() {
	fmt.Println("consumer1 run...")
	simple := mq集合.NewRabbitMQSimple("kuteng")

	simple.ConsumeSimple("consumer1")
}

func consumer2() {
	fmt.Println("consumer2 run..")
	simple := mq集合.NewRabbitMQSimple("kuteng")

	simple.ConsumeSimple("consumer2")
}

// 简单的消费者
func main() {

	go consumer1()
	go consumer2()
	for true {
		time.Sleep(100 * time.Second)
	}

}
