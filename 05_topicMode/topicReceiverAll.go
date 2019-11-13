package main

import "go-rabbitmq/00_RabbitMQ"

func main() {
	/*
		匹配所有
	*/
	rabbitMQRoutingAll := _0_RabbitMQ.NewRabbitMQRouting(_0_RabbitMQ.TopicExchangeName, "#")

	rabbitMQRoutingAll.ReceiverTopic()
}
