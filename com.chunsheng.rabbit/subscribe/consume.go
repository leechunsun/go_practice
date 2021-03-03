package main

import "com.chunsheng.test/com.chunsheng.rabbit/base"


func main() {
	exchange := "exchange20201009"
	mq := base.NewSubscribe(exchange)
	mq.SubConsume()
}
