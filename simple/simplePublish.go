package main

import (
	"go-rabbitmq/RabbitMQ"
	"log"
)

func main() {
	rabbitMQSimple := RabbitMQ.NewRabbitMQSimple(RabbitMQ.SimpleQueueName)
	rabbitMQSimple.PublishSimple("hello 简单模式消息")
	log.Println("发送成功")
}
