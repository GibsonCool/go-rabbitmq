package main

import "go-rabbitmq/00_RabbitMQ"

func main() {

	/*
		仅仅匹配 "double.*.two" 比如我们生产端的 double.topic.two
	*/
	rabbitMQRoutingTwo := _0_RabbitMQ.NewRabbitMQRouting(_0_RabbitMQ.TopicExchangeName, "double.*.two")

	rabbitMQRoutingTwo.ReceiverTopic()
}
