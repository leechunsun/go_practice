package base

import (
	"fmt"
	"github.com/streadway/amqp"
)

const MQURL = "amqp://guest:guest@127.0.0.1:5672/simple"


type MyMQ struct {
	Conn *amqp.Connection
	Channel *amqp.Channel
	QueueName string
	Exchange string
	Key string
	Url string
}


func NewMYMQ(queueName, exchange, key string) *MyMQ {
	return &MyMQ{QueueName:queueName, Exchange:exchange, Key:key, Url:MQURL}
}


func (m *MyMQ) Destory(){
	m.Channel.Close()
	m.Conn.Close()
}


func (m *MyMQ) FailOnErr(err error, message string) {
	if err != nil {
		fmt.Printf("err是:%s,信息是:%s", err, message)
}
}





