package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Sharek(origin []int) {
	l := len(origin)
	for i:=0;i < l;i++{
		seed := time.Now().UnixNano()
		x := rand.New(rand.NewSource(seed)).Intn(l)
		origin[i], origin[x] = origin[x], origin[i]
	}
}

func main() {
	x := []int{0,1,2,3,4,5,6,7,8,9}
	for k:=0;k<10;k++{
		Sharek(x)
		fmt.Println(k," : ",x)
	}
}


