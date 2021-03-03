package main

import "com.chunsheng.test/com.chunsheng.rabbit/base"



func main() {
	var SimpleQueueName = "chunsheng20200930"
	mq := base.NewSimpleMQ(SimpleQueueName)
	mq.SimpleConsume()
}