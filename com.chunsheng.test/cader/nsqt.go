package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

var productor *nsq.Producer
var consumer *nsq.Consumer


type C struct {

}

func (c C) HandleMessage(msg *nsq.Message) error {
	fmt.Println("================================")
	fmt.Println(string(msg.Body))
	return nil
}


func init(){
	var err error
	productor, err = nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())
	if err != nil{
		fmt.Println("dail err :", err)
		return
	}
	consumer, err = nsq.NewConsumer("test", "tc", nsq.NewConfig())
	if err != nil{
		fmt.Println("dail err :", err)
		return
	}
	consumer.AddHandler(C{})
	consumer.ConnectToNSQLookupd("127.0.0.1:4161")

}

func main() {
	fmt.Print(".....")

	for {
		time.Sleep(3 * time.Second)
		productor.Publish("test", []byte(fmt.Sprintf("现在的时间是：%d", time.Now().Unix())))
	}
}



