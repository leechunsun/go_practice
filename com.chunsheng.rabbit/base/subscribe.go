package base

import (
	"fmt"
	"github.com/streadway/amqp"
)

func NewSubscribe(exchangeName string) *MyMQ {
	mq := &MyMQ{QueueName:"", Exchange: exchangeName, Key:""}
	var err error
	mq.Conn, err = amqp.Dial(MQURL)
	if err != nil{
		fmt.Println("连接rabbitmq 失败")
		panic(err)
	}
	mq.Channel, err = mq.Conn.Channel()
	if err != nil{
		fmt.Println("获取rabbitmq channel 失败")
		panic(err)
	}
	fmt.Println("创建订阅模式。。。。")
	return mq
}

func (m *MyMQ) SubPublish(msg string){
	err := m.Channel.ExchangeDeclare(m.Exchange,
										"fanout",
										true,
										false,
										false,
										false,
										nil,
										)
	m.FailOnErr(err, "交换机定义失败")
	m.Channel.Publish(m.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			Type:            "text/plain",
			Body:            []byte(msg),
		})
}


func (m *MyMQ) SubConsume(){
	err := m.Channel.ExchangeDeclare(m.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	m.FailOnErr(err, "交换机定义失败")
	q, err := m.Channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil)
	m.FailOnErr(err, "订阅模式消费方法中创建队列失败。")

	err = m.Channel.QueueBind(q.Name,"", m.Exchange, false, nil)
	m.FailOnErr(err, "订阅模式消费方法中绑定队列失败。")
	msg, err := m.Channel.Consume(q.Name, "", true, false, false, false, nil)
	for x := range msg{
		fmt.Println("收到消息：", string(x.Body))
	}
}

