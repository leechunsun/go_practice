package main

import "com.chunsheng.test/com.chunsheng.rabbit/base"

func main() {
	exchange := "20201009/routing"
	mq2 := base.NewRoutingMq(exchange, "Queue2")
	mq2.RoutingConsume()
}
