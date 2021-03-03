package main

import (
	"fmt"
	"time"
)

func main() {
	z := make(chan string)
	count := 3
	for i := 0;i< count; i++{
		x := i
		go func() {
			time.Sleep(time.Duration(2000 * x)*time.Millisecond)
			fmt.Println("hello ", <-z)
		}()
	}

	for it:=0;it<count;it++ {
		z <- fmt.Sprintf("i am %d", it)
	}
	time.Sleep(time.Duration(2000 * 5)*time.Millisecond)
}
