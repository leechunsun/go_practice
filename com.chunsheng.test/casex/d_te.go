package main

import (
	"context"
	"fmt"
	"time"
)

type Myc struct {
	Name string `config:"name"`
	Age int `config:"age"`
	Addr string `config:"addr"`
}

func (m *Myc) callTest(){
	fmt.Println(time.NewTicker(1 * time.Second))
}


func ZZ() interface{} {
	ctx := context.WithValue(context.Background(), "zz", "kkkk")
	return ctx.Value("zz")

}

func main() {
	message := make(chan int, 10)

	for j:=0;j<10;j++{
		message <- j
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	go func(cont context.Context) {
		tiker := time.NewTicker(time.Second)
		for _ = range tiker.C{
			select {
			case <- ctx.Done():
				fmt.Println("ctx be canceled....")
			default:
				fmt.Printf("get_message: %d", <-message)
			}
		}
	}(ctx)

	defer close(message)
	defer cancel()
	select {
	case <-ctx.Done():
		fmt.Println("program exiting....")
	}

}