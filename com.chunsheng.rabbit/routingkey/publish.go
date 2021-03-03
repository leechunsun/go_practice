package main

import (
	"com.chunsheng.test/com.chunsheng.rabbit/base"
	"strconv"
	"time"
)


func main() {
	exchange := "20201009/routing"
	mq1 := base.NewRoutingMq(exchange, "Queue1")
	mq2 := base.NewRoutingMq(exchange, "Queue2")

	for i:=0 ; i<100; i++{
		msg := "gan!!! routing---msg" + strconv.Itoa(i)
		if i % 2 == 0 {
			mq1.RoutingPublish(msg)
		}else{
			mq2.RoutingPublish(msg)
		}
		time.Sleep(time.Second * 1)
	}
}
