package main

import (
	"fmt"
	_0_RabbitMQ "go-rabbitmq/00_RabbitMQ"
	"log"
	"time"
)

/*
	路由模式：
		一个生产端，多个消费端。路由模式下，生产端可以根据路由 Key 指定那种路由 key 消费端接收消费消息
*/
func main() {
	rabbitMQRoutingOne := _0_RabbitMQ.NewRabbitMQRouting(_0_RabbitMQ.RoutingExchangeName, _0_RabbitMQ.RoutingKey1)
	rabbitMQRoutingTwo := _0_RabbitMQ.NewRabbitMQRouting(_0_RabbitMQ.RoutingExchangeName, _0_RabbitMQ.RoutingKey2)
	for i := 0; i <= 10; i++ {
		log.Printf("正在发送第 %d 条消息~~~", i)
		rabbitMQRoutingOne.PublishRouting(fmt.Sprintf("Routing key:%s 路由模式消息 index:%d", rabbitMQRoutingOne.Key, i))
		if i%2 == 0 {
			rabbitMQRoutingTwo.PublishRouting(fmt.Sprintf("Routing key:%s 路由模式消息 index:%d", rabbitMQRoutingTwo.Key, i))
		}
		time.Sleep(1 * time.Second)
	}

}
