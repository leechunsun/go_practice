package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

const (
	CallingType = 0
	CalledType = 1
)


// 写入文件逻辑
func Write(condtionArr []int, writeChan chan string, strType int, MidDirPath string){
	bufmap := sync.Map{}
	ch := make(chan bool, WriteMaxGo)
	wc := sync.Map{}
	for str := range writeChan{
		ch <- true
		go func(str string) {
			defer func() {
				if err:= recover(); err != nil{

				}
			}()
			var strcall int
			var strin string
			var err error
			var str_called string
			if strType == CallingType{
				sind := strings.Index(str, ",")
				str_pre := str[:sind]
				strcall, err = strconv.Atoi(strings.Trim(str_pre, " "))
				strin = str
			}else if strType == CalledType{
				allins := strings.Split(str, ",")
				if len(allins) > 2{
					strCalledLen := len(allins[1])
					if strCalledLen >= 9 {
						str_called = strings.Trim(allins[1], "")[strCalledLen-9:]
					}else{
						<- ch
						return
					}
					strcall, err = strconv.Atoi(strings.Trim(str_called, " "))
					allins[0], allins[1] = str_called, allins[0]
					strin = strings.Join(allins, ",")
				} else {
					<- ch
					return
				}
			}
			if err != nil{
				<- ch
				return
			}
			fileCase := SideFind(condtionArr, strcall)[0]
			fileP := filepath.Join(MidDirPath, strconv.Itoa(fileCase))
			if loc, _ := wc.LoadOrStore(fileP, make(chan bool, 1)); true{
				loc.(chan bool) <- true
				if val, _ := bufmap.LoadOrStore(fileP, []string{});true{
					valArr := append(val.([]string), strin + "\n")
					if len(valArr) >= 3000{
						go func(valArr []string) {
							f, err := os.OpenFile(fileP, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
							if err != nil{
								fmt.Println("Error 写入文件失败：", str)
							}
							buff := bufio.NewWriter(f)
							for _, val := range valArr{
								buff.WriteString(val)
							}
							buff.Flush()
							defer f.Close()
						}(valArr)
						valArr = []string{}
					}
					bufmap.Store(fileP, valArr)
				}
				<- loc.(chan bool)
			}
			<- ch
		}(str)
	}
	for _, fileName := range condtionArr{
		fileCase := filepath.Join(MidDirPath, strconv.Itoa(fileName))
		if valArr, ok := bufmap.Load(fileCase);ok{
			f, err := os.OpenFile(fileCase, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil{
				fmt.Println("Error 刷新文件失败：", err)
			}
			buff := bufio.NewWriter(f)
			for _, val := range valArr.([]string){
				buff.WriteString(val)
			}
			buff.Flush()
			f.Close()
		}
	}
	fmt.Println("完成写入......")
}

func Dispatch(writeChan chan string, filePath string) {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return
	}
	buf := bufio.NewReader(f)
	cond := 0
	for {
		if cond == 0{
			cond = 1
		}
		line, _ , err := buf.ReadLine()
		oneData := strings.TrimSpace(string(line))
		writeChan <- oneData
		if err != nil {
			if err == io.EOF{
				return
			}
			return
		}
	}
}






