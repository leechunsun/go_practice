package main

import (
	"fmt"
	"sync"
	"time"
)

var pool sync.Pool


func main() {
	c := make(chan bool)
	pool.Put("zzzzzzz")
	go func() {
		time.Sleep(time.Second * 3)
		<- c
		fmt.Println(pool.Get())
	}()
	c <- false
	fmt.Println("kkkk")
}
