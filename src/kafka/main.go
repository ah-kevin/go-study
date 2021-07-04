package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	// 1. 生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // Ack发完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition // 分区
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回 确认

	// 2. 连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed,err:", err)
	}
	defer func(client sarama.SyncProducer) {
		err := client.Close()
		if err != nil {
			fmt.Println("producer closed,err:", err)
		}
	}(client)

	// 3. 封装消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "shopping"
	msg.Value = sarama.StringEncoder("哈哈哈哈哈")

	// 4. 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed,err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
