package main

import (
	"fmt"
	"localcache/db_instance"
	"localcache/http_handler"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	ars := os.Args
	port := getPorts(ars)
	db := db_instance.GetCacheDBSingle()
	go reverseCheck(db)
	stringHandler := &http_handler.StringHandler{DB:db}
	http.Handle("/string", stringHandler)
	lInfos(port)
	err := http.ListenAndServe(port, nil)
	if err != nil{
		fmt.Println("l cache system stop by err:", err)
	}
}

func reverseCheck(db *db_instance.CacheDB){
	for{
		time.Sleep(time.Minute * 1)
		db.CheckOverline()
	}
}

func lInfos (port string){
	fmt.Println("# l cache system started.")
	fmt.Println("# listen-", port)
	fmt.Println("# version v1.0")
	fmt.Println("#", time.Now())
}

func getPorts(a []string) string {
	var port = ":7777"
	if len(a) >= 2{
		port = a[1]
	}else{
		return port
	}
	if d, err := strconv.Atoi(port); err == nil {
		return fmt.Sprintf(":%d", d)
	}
	return port
}
