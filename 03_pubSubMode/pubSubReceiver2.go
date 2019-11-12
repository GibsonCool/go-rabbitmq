package main

import _0_RabbitMQ "go-rabbitmq/00_RabbitMQ"

func main() {

	rabbitMQPubSub := _0_RabbitMQ.NewRabbitMQPubSub(_0_RabbitMQ.PubSubExchangeName)
	rabbitMQPubSub.ReceiverSub()
}
