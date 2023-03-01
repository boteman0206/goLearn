package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
)

/**
kafka代码演示消费者端  消费者组的概念
*/

func NewConsumerGroup(ctx context.Context, groupName string) (sarama.ConsumerGroup, error) {

	config := sarama.NewConfig()

	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	return sarama.NewConsumerGroup([]string{"127.0.0.1:9092"}, groupName, config)
}

func main() {
	group, err := NewConsumerGroup(context.Background(), "group1")
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	if err != nil {
		fmt.Println("kafkaLib.NewConsumerGroup failed with err: %s", err.Error())
		return
	}
	defer group.Close()

	for {
		err = group.Consume(context.Background(), []string{"test"}, msgHandler{})
		if err != nil {
			fmt.Println("group.Consume failed with err: %s", err.Error())
		}
	}

	select {}
}

/*
*
实现消费者组方法
*/
type msgHandler struct{}

// Setup is run at the beginning of a new session, before ConsumeClaim.
func (msgHandler) Setup(session sarama.ConsumerGroupSession) error { return nil }

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (msgHandler) Cleanup(session sarama.ConsumerGroupSession) error { return nil }

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (handler msgHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		if err := dispatchMsg(context.Background(), msg); err != nil {
			fmt.Println("dispatch_msg_failed err: %s, msg topic:%s, partition:%d, offset:%d, value: %s",
				err.Error(),
				msg.Topic,
				msg.Partition,
				msg.Offset,
				string(msg.Value),
			)
		} else {
			// 如果处理消息成功，则提交消费位移
			session.MarkMessage(msg, "")
		}
	}
	return nil
}

func dispatchMsg(ctx context.Context, msg *sarama.ConsumerMessage) error {
	fmt.Println(string(msg.Value))

	return nil
}
