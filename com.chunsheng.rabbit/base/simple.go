package base

import (
	"fmt"
	"github.com/streadway/amqp"
)

func NewSimpleMQ(QueueName string) *MyMQ {
	rabbitmq := NewMYMQ(QueueName, "", "")
	var err error
	rabbitmq.Conn, err = amqp.Dial(rabbitmq.Url)
	rabbitmq.FailOnErr(err, "连接connection失败")

	rabbitmq.Channel, err = rabbitmq.Conn.Channel()
	rabbitmq.FailOnErr(err, "获取channel参数失败")
	return rabbitmq
}

func (m *MyMQ) PublishSimple(msg string){
	// 创建队列
	_, err := m.Channel.QueueDeclare(m.QueueName,
		false,
		false,
		false,
		false,
		nil)
	m.FailOnErr(err, "创建队列失败")

	m.Channel.Publish(m.Exchange,
		m.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
}


func (m *MyMQ) SimpleConsume(){
	_, err := m.Channel.QueueDeclare(m.QueueName,
		false,
		false,
		false,
		false,
		nil)
	if err != nil{
		fmt.Println(err)
	}

	msg, err :=m.Channel.Consume(m.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil)
	if err != nil{
		fmt.Println(err)
	}

	for m := range msg{
		fmt.Println("接受消息:", string(m.Body))
	}
}


