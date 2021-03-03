package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(x int) {
			defer wg.Done()
			fmt.Println(x)
		}(i)
	}
	wg.Wait()
}
