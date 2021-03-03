package main

import (
	"com.chunsheng.test/com.chunsheng.rabbit/base"
	"strconv"
)

func main() {
	var queueName = "work20200930"
	mq := base.NewSimpleMQ(queueName)
	for i:=0;i <= 100;i++{
		mq.PublishSimple("i am " + strconv.Itoa(i))
	}
}
