package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	MaxGoPRO = 8
	WriteQueueSize = 10000
	WriteMaxGo = 200
	MaxFileDeal = 10
)



func main(){
	fmt.Println("===========================")
	fmt.Println("打开方式: ./运行文件 源文件路径 目标文件路径 处理类型（0:calling, 1:called）")
	fmt.Println("===========================")
	runtime.GOMAXPROCS(MaxGoPRO)
	var (
		DirPath = "/data1/work/lcs/test"
		MidDirPath = "/data1/work/lcs/testmid"
		DealType = "0"
	)
	ars := os.Args
	if len(ars)>0 {
		if len(ars) >= 2 {
			DirPath = ars[1]
		}
		if len(ars) >= 3 {
			MidDirPath = ars[2]
		}
		if len(ars) >= 4 {
			DealType = ars[3]
		}
	}
	fmt.Println("源文件路径：", DirPath)
	fmt.Println("目标文件路径：", MidDirPath)
	fmt.Println("处理文件类型：", DealType)
	dType, err := strconv.Atoi(DealType)
	if err != nil || (dType != CallingType && dType != CalledType){
		panic("请输入正确的处理类型.......")
	}
	wg := sync.WaitGroup{}
	filesInfo, err := ioutil.ReadDir(DirPath)
	if err != nil {
		panic(err)
	}
	var condArr []int
	var FileNameArr []string
	cond := 0
	for _, fileInfo := range filesInfo{
		if ! strings.HasSuffix(fileInfo.Name(), "csv"){
			continue
		}
		if cond == 0{
			condArr = InitCondArr(filepath.Join(DirPath, fileInfo.Name()))
			cond = 1
		}
		FileNameArr = append(FileNameArr, fileInfo.Name())
	}
	tall := time.Now().Unix()
	// 定义写文件的channel 并运行循环逻辑
	writerChannel := make(chan string, WriteQueueSize)
	// 执行处理文件逻辑
	go Write(condArr, writerChannel, dType, MidDirPath)
	fmt.Println("开始处理文件, 文件个数：", len(FileNameArr))
	// 进行文件分发工作
	chf := make(chan bool, MaxFileDeal)
	for _, fileName := range FileNameArr{
		if ! strings.HasSuffix(fileName, "csv"){
			continue
		}
		chf <- true
		wg.Add(1)
		go func(fileNamea string) {
			t := time.Now().Unix()
			Dispatch(writerChannel, filepath.Join(DirPath, fileNamea))
			fmt.Println("分发完成：", fileNamea, "， 耗时：", time.Now().Unix()-t, "s....")
			<- chf
			wg.Done()
		}(fileName)
	}
	wg.Wait()
	close(writerChannel)
	fmt.Println("分发结束， 耗时：", time.Now().Unix()-tall, "s....")
	time.Sleep(3 * time.Second)
}
