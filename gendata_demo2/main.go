package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

const (
	MaxGoPRO = 8
	MaxFileDeal = 1
	WriteChanSize = 10000
)

var WChan = make(chan *MyFile, WriteChanSize)

var wg = sync.WaitGroup{}


func main() {
	fmt.Println("===========================")
	fmt.Println("打开方式: ./运行文件 month 源文件路径 目标文件路径 文件格式")
	fmt.Println("===========================")
	runtime.GOMAXPROCS(MaxGoPRO)
	var (
		DirPath = "/data1/work/lcs/testmid"
		FinDirPath = "/data1/work/lcs/testfin"
		fileType = "mysql"
		month = ""
	)
	ars := os.Args
	if len(ars)>0 {
		if len(ars) >= 2 {
			month = ars[1]
		}
		if len(ars) >= 3 {
			DirPath = ars[2]
		}
		if len(ars) >= 4 {
			FinDirPath = ars[3]
		}
		if len(ars) >= 5 {
			fileType = ars[4]
		}
	}
	if month == ""{
		panic("请输入月份")
	}
	fmt.Println("源文件路径：", DirPath)
	fmt.Println("目标文件路径：", FinDirPath)
	fmt.Println("写入格式：", fileType)
	fmt.Println("月份信息：", month)
	fileInfos, err := ioutil.ReadDir(DirPath)
	if err != nil {
		panic(err)
	}
	var fileNames []string
	for _, fileInfo := range fileInfos {
		fileNames = append(fileNames, fileInfo.Name())
	}
	t1 := time.Now().Unix()
	proChan := make(chan bool, MaxFileDeal)
	wg.Add(1)
	go func() {
		Write()
		wg.Done()
	}()
	var realFileName string
	for _, fileNameo := range fileNames{
		proChan <- true
		wg.Add(1)
		go func(fileName string) {
			t := time.Now().Unix()
			switch fileType {
			case "mysql":
				realFileName = fileName + ".csv"
			case "json":
				realFileName = fileName + ".json"
			default:
				panic(fileType + "格式不支持....")
			}
			origin := filepath.Join(DirPath, fileName)
			filePath := filepath.Join(FinDirPath, realFileName)
			Gen(origin, filePath, fileType, month)
			fmt.Println("文件：", fileName, ", 写入完成, 耗时：", time.Now().Unix()-t, "s....")
			<- proChan
			wg.Done()
		}(fileNameo)
	}
	wg.Wait()
	fmt.Println("文件写入完成.....", time.Now().Unix()-t1, "s")
}


