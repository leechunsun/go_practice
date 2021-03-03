package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func InitCondArr(filepath string) []int{
	t := time.Now().Unix()
	fmt.Println("开始初始化切片列表......")
	var condArr []int
	var midArr []int
	f, err := os.Open(filepath)
	if err != nil{
		panic(err)
	}
	defer f.Close()
	bufreader := bufio.NewReader(f)
	for i := 0; i < 2000000;i++{
		info, _, err := bufreader.ReadLine()
		if i == 0 || err != nil{
			continue
		}
		sind := strings.Index(string(info), ",")
		cas, err:= strconv.Atoi(strings.Trim(string(info)[:sind], " "))
		if err != nil{
			continue
		}
		midArr = append(midArr, cas)
	}
	sort.Ints(midArr)
	for i, val := range midArr{
		if i % 4000 == 0 && i / 4000 > 0{
			condArr = append(condArr, val)
		}
	}
	fmt.Println("完成初始化切片列表...", "耗时:", time.Now().Unix()-t, "s...")
	return condArr
}
