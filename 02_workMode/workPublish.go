package main

import (
	"fmt"
	"go-rabbitmq/00_RabbitMQ"
	"log"
)

/*
	worker 模式：
		生产速度大于消费速度，一个生产端，多个消费端。和simple模式一样，一个消息只能被一个消费端消费

		workReceiver1  workReceiver2 workReceiver2 都是一模一样的代码，起到的作用其实就是负载均衡的效果
		从运行的结果可以看出这里的负载均衡策略基本是轮询三个消费端接收
*/
func main() {
	rabbitMQSimple := _0_RabbitMQ.NewRabbitMQSimple(_0_RabbitMQ.SimpleQueueName)
	for i := 0; i <= 100; i++ {
		log.Printf("正在发送第 %d 条消息~~~", i)
		rabbitMQSimple.PublishSimple(fmt.Sprintf("hello 简单模式消息 index:%d", i))
	}

}
