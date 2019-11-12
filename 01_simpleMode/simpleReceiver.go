package main

import "go-rabbitmq/00_RabbitMQ"

func main() {
	rabbitMQSimple := _0_RabbitMQ.NewRabbitMQSimple(_0_RabbitMQ.SimpleQueueName)

	rabbitMQSimple.ConsumeSimple()
}
