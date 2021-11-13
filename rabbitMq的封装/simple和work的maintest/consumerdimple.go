package main

import (
	"fmt"
	"godemo/rabbitMq的封装/simple和work工作模式mq集合"
	"time"
)

func consumer1() {
	fmt.Println("consumer1 run...")
	simple := simple和work工作模式mq集合.NewRabbitMQSimple("kuteng")

	simple.ConsumeSimple("consumer1")
}

func consumer2() {
	fmt.Println("consumer2 run..")
	simple := simple和work工作模式mq集合.NewRabbitMQSimple("kuteng")

	simple.ConsumeSimple("consumer2")
}

// 简单的消费者    todo simple工作模式和work工作模式  work工作模式只是在simple的基础上多了多个消费者
func main() {

	go consumer1() // 消费者1
	go consumer2() // 消费者2
	for true {
		time.Sleep(100 * time.Second)
	}

}
