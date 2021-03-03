package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"syscall"
)

func LastCoins(coins []int, digest int) int {
	all_money := make([]int, digest+1)
	for _, money := range coins{
		for ari, c := range all_money{
			if ari >= money && (all_money[ari - money] > 0 || ari == money){
				need := all_money[ari - money] + 1
				if need < c || all_money[ari] == 0{
					all_money[ari] = need
				}
			}
		}
	}
	fmt.Println(all_money)
	return all_money[digest]
}

func main() {
	coins := []int{1,2,5}
	fmt.Println(LastCoins(coins, 50))
	req, err := http.Get("http://sina.com")
	if err != nil{
		fmt.Println(err)
	}

	defer req.Body.Close()
	x, err := ioutil.ReadAll(req.Body)
	fmt.Println(string(x), err)
}

