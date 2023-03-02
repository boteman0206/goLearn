package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

/*
*

	代码演示  异步发送
*/
func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	config.Producer.Return.Errors = true                      // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "test"
	msg.Value = sarama.StringEncoder("async send message 01")
	// 连接kafka
	// 异步发送
	clientAsync, err := sarama.NewAsyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	// send to chain
	clientAsync.Input() <- msg
	select {
	case suc := <-clientAsync.Successes():
		fmt.Println("[kafka-send]: producer send msg success, partition: %d, msg: %+v", suc.Partition, msg)
	case fail := <-clientAsync.Errors():
		fmt.Println("[kafka-send]: producer send msg failed, err: %s, msg: %+v", fail.Error(), msg)
	}

	fmt.Printf("async success end！！！")

}
