package main

import (
	"go-rabbitmq/00_RabbitMQ"
	"log"
	"strconv"
	"time"
)

/*
	订阅模式：
		一个生产端，多个消费端。和worker模式不通的是，只要订阅了生产端的消费端，由消费端产生的消息每个
		消费端都能接收到。有点类似于生产端是广播的形式通过 交换机将消息发送到各个消费端对应绑定的队列中
*/
func main() {
	rabbitMQPubSub := _0_RabbitMQ.NewRabbitMQPubSub(_0_RabbitMQ.PubSubExchangeName)

	for i := 0; i < 10; i++ {
		msg := "订阅模式生产index：" + strconv.Itoa(i) + "数据"
		rabbitMQPubSub.PublishPub(msg)
		log.Println(msg)
		time.Sleep(1 * time.Second)
	}
}
