package main

import (
	"fmt"
	_0_RabbitMQ "go-rabbitmq/00_RabbitMQ"
	"log"
	"time"
)

/*
	Topic 模式下：
		routing_key:必须是由点隔开的一系列的标识符组成。标识符可以是任何东西，但是一般都与消息的某些特性相关

		*可以匹配一个标识符。
		#可以匹配0个或多个标识符
*/
func main() {
	rabbitMQRoutingOne := _0_RabbitMQ.NewRabbitMQRouting(_0_RabbitMQ.TopicExchangeName, _0_RabbitMQ.TopicKey1)
	rabbitMQRoutingTwo := _0_RabbitMQ.NewRabbitMQRouting(_0_RabbitMQ.TopicExchangeName, _0_RabbitMQ.TopicKey2)
	for i := 0; i <= 10; i++ {
		log.Printf("正在发送第 %d 条消息~~~", i)
		rabbitMQRoutingOne.PublishTopic(fmt.Sprintf("Routing key:%s topic模式消息 index:%d", rabbitMQRoutingOne.Key, i))
		rabbitMQRoutingTwo.PublishTopic(fmt.Sprintf("Routing key:%s topic模式消息 index:%d", rabbitMQRoutingTwo.Key, i))
		time.Sleep(1 * time.Second)
	}

}
