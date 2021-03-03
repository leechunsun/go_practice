package main

import (
	"com.chunsheng.test/set"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	ss := &set.StringSet{}
	t := time.Now().UnixNano()
	for i:=0; i<1000000;i++{
		wg.Add(1)
		go func(i int) {
			//ss.Add(strconv.Itoa(i))
			ss.Contains(strconv.Itoa(i))
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Printf("spent: %.3f s", float64(time.Now().UnixNano() - t) / (1000 * 1000 * 1000) )
	fmt.Println()
	fmt.Println(ss.Len())
}
