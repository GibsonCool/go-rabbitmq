package main

import "go-rabbitmq/RabbitMQ"

func main() {
	rabbitMQSimple := RabbitMQ.NewRabbitMQSimple(RabbitMQ.SimpleQueueName)

	rabbitMQSimple.ConsumeSimple();
}
