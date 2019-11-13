package main

import "go-rabbitmq/00_RabbitMQ"

func main() {
	/*
		只能接受到 RoutingKey1 的消息
	*/
	rabbitMQRoutingTwo := _0_RabbitMQ.NewRabbitMQRouting(_0_RabbitMQ.RoutingExchangeName, _0_RabbitMQ.RoutingKey1)

	rabbitMQRoutingTwo.ReceiverRouting()
}
