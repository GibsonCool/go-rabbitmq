package main

import "go-rabbitmq/00_RabbitMQ"

func main() {
	/*
		无法接收到任何消息，没有对应的  key
	*/
	rabbitMQRoutingOne := _0_RabbitMQ.NewRabbitMQRouting(_0_RabbitMQ.RoutingExchangeName, "null")

	rabbitMQRoutingOne.ReceiverRouting()
}
