package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// url 格式  amqp://账号:密码@rabbitmq服务器地址:端口号/vhost
const (
	MQURL           = "amqp://double:double@127.0.0.1:5672/test-double"
	SimpleQueueName = "doubleSimple"
)

var Done chan bool

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// binding key
	Key string
	// 连接信息
	Mqurl string
}

// 创建 RabbitMQ 实例
func NewRabbitMq(queueName string, exchange string, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
	// 创建 rabbitmq 链接
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnError(err, "创建连接错误！")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnError(err, "获取 channel 失败")
	return rabbitmq
}

// 断开 channel/connection
func (r *RabbitMQ) Destory() {
	_ = r.conn.Close()
	_ = r.channel.Close()
}

// 错误处理函数
func (r *RabbitMQ) failOnError(err error, message string) {
	if err != nil {
		log.Printf("%s:%v", message, err)
		panic(fmt.Sprintf("%s:%v", message, err))
	}
}

// 简单模式 Step1：创建简单模式下 RabbitMQ 实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	//简单模式交换机使用的是默认的
	return NewRabbitMq(queueName, "", "")
}

// 简单模式 Step2：简单模式下生产代码
func (r *RabbitMQ) PublishSimple(message string) {
	// 1.申请队列，如果队列不存在则会自动创建，如果存在则跳过创建
	// 好处：保证队列存在，消息能发送到队列中
	_, e := r.channel.QueueDeclare(
		r.QueueName,
		// 是否持久化
		false,
		// 是否会自动删除：当最后一个消费者断开后，是否将队列中的消息清除
		false,
		// 是否具有排他性，意思只有自己可见，其他人不可用
		false,
		// 是否阻塞
		false,
		// 其他额外属性
		nil,
	)
	if e != nil {
		fmt.Println(e)
	}

	// 2.发送消息到队列中
	e = r.channel.Publish(
		r.Exchange,
		r.QueueName,
		// 如果是 true，根据 exchange 类型 routkey 规则，如果无法找得到符合条件的队列那么会把消息返回给发送者
		false,
		// 如果是 true，当 exchange 发送消息到队列后发现队列上没有绑定消费者，会将消息返送回给发送者
		false,
		// 真正的消息
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if e != nil {
		fmt.Println(e)
	}

}

// 简单模式 Step3：简单模式下消费代码
func (r *RabbitMQ) ConsumeSimple() {
	// 无论生产还是消费，第一步都是尝试先申请队列
	_, e := r.channel.QueueDeclare(
		r.QueueName,
		// 是否持久化
		false,
		// 是否会自动删除：当最后一个消费者断开后，是否将队列中的消息清除
		false,
		// 是否具有排他性，意思只有自己可见，其他人不可用
		false,
		// 是否阻塞
		false,
		// 其他额外属性
		nil,
	)
	if e != nil {
		fmt.Println(e)
	}

	Done = make(chan bool)
	msgs, e := r.channel.Consume(
		r.QueueName,
		//用来区分多个消费者，消费者处理器名称
		"",
		//是否自动应答通知已收到消息
		true,
		//是否排他性,非唯一的消费者，其他消费者处理器也可以去竞争这个队列里面的消息任务
		false,
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false,
		//是否阻塞
		false,
		nil,
	)

	go func() {
		for d := range msgs {
			// 这里接收到消息，实现处理逻辑
			log.Printf("接收到消息：%s", d.Body)
		}
	}()

	log.Printf("消费者已开启，等待消息产生。。。")
	<-Done

	r.Destory()
	log.Printf("消费者关闭。。。")
}

// 订阅模式创建 RabbitMQ 实例
func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
	// 创建 RabbitMQ 实例
	rabbitMQ := NewRabbitMq("", exchangeName, "")

}
