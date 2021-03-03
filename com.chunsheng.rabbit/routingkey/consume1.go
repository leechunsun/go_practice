package main

import "com.chunsheng.test/com.chunsheng.rabbit/base"

func main() {
	exchange := "20201009/routing"
	mq1 := base.NewRoutingMq(exchange, "Queue1")
	mq1.RoutingConsume()
}
