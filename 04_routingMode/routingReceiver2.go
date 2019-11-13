package main

import "go-rabbitmq/00_RabbitMQ"

func main() {
	/*
		只能接受到 RoutingKey2 的消息
	*/
	rabbitMQRoutingOne := _0_RabbitMQ.NewRabbitMQRouting(_0_RabbitMQ.RoutingExchangeName, _0_RabbitMQ.RoutingKey2)

	rabbitMQRoutingOne.ReceiverRouting()
}
