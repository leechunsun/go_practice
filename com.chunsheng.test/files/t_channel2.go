package main

import (
	"fmt"
	"time"
)

func main() {
	v := make(chan string)
	c := 3
	for x := 1;x < c; x++ {
		z := x
		go func() {
			time.Sleep(time.Duration(2000 * z)* time.Millisecond)
			v <- fmt.Sprintf("hello iam %d", z)
		}()
	}

	for i := 0;i < c + 3; i++ {
		fmt.Println(<-v)
	}
}
