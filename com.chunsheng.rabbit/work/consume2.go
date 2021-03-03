package main

import "com.chunsheng.test/com.chunsheng.rabbit/base"

func main() {
	var queueName = "work20200930"
	mq := base.NewSimpleMQ(queueName)
	mq.SimpleConsume()
}