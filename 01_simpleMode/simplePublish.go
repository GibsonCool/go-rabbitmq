package main

import (
	"go-rabbitmq/00_RabbitMQ"
	"log"
)

func main() {
	rabbitMQSimple := _0_RabbitMQ.NewRabbitMQSimple(_0_RabbitMQ.SimpleQueueName)
	rabbitMQSimple.PublishSimple("hello 简单模式消息")
	log.Println("发送成功")
}
