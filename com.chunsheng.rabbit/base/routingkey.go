package base

import (
	"fmt"
	"github.com/streadway/amqp"
)

func NewRoutingMq(exchange, routingkey string) *MyMQ {
	mq := &MyMQ{QueueName: "", Exchange:exchange, Key:routingkey}
	var err error
	mq.Conn, err = amqp.Dial(MQURL)
	mq.FailOnErr(err, "mq连接失败")
	mq.Channel, err = mq.Conn.Channel()
	mq.FailOnErr(err, "mq连接channel失败")
	fmt.Println("routing key 模式。。。。")
	return mq
}


func (m *MyMQ) RoutingPublish(msg string) {
	err := m.Channel.ExchangeDeclare(m.Exchange, "direct", false, false, false, false, nil)
	m.FailOnErr(err, "尝试创建交换机失败")
	q, err := m.Channel.QueueDeclare("", false, false, false, false, nil)
	m.FailOnErr(err, "尝试创建队列失败")
	err = m.Channel.QueueBind(q.Name, m.Key, m.Exchange, false, nil)
	m.FailOnErr(err, "尝试绑定失败")
	m.Channel.Publish(m.Exchange, m.Key, false, false, amqp.Publishing{
		Type:            "text/plain",
		Body:            []byte(msg),
	})
}


func (m *MyMQ) RoutingConsume(){
	err := m.Channel.ExchangeDeclare(m.Exchange, "direct", false, false, false, false, nil)
	m.FailOnErr(err, "尝试创建交换机失败")
	q, err := m.Channel.QueueDeclare("", false, false, false, false, nil)
	m.FailOnErr(err, "尝试创建队列失败")
	err = m.Channel.QueueBind(q.Name, m.Key, m.Exchange, false, nil)
	m.FailOnErr(err, "尝试绑定失败")
	msg, err := m.Channel.Consume(m.QueueName, "", true, false, false, false, nil)
	for x := range msg {
		fmt.Println("routingKey:", m.Key, "  收到消息：", string(x.Body))
	}
}


