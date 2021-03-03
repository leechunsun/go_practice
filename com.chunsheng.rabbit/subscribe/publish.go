package main

import (
	"com.chunsheng.test/com.chunsheng.rabbit/base"
	"strconv"
)

func main() {
	exchange := "exchange20201009"
	mq := base.NewSubscribe(exchange)
	for i:=0;i<100;i++{
		mq.SubPublish("exchange-" + strconv.Itoa(i)+": gan!")
	}
}





