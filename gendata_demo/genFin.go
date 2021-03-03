package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GenMysql(wf *os.File,finMap *map[string]*strings.Builder, month string) {
	wb := bufio.NewWriter(wf)
	cnt := 0
	for key, value := range *finMap{
		realKey_int, err :=  Cast89En(key)
		realKey := strconv.Itoa(realKey_int)
		if err !=  nil{
			continue
		}
		valStr := ZipBase64(value.String())
		wb.WriteString(realKey + "," + month + "," + valStr + "\n")
		cnt ++
		if cnt > 999 && cnt % 1000 == 0{
			wb.Flush()
		}
	}
	wb.Flush()
	wf.Close()
	wg.Done()
}


func GenJson(wf *os.File,finMap *map[string]*strings.Builder, month string) {
	wb := bufio.NewWriter(wf)
	cnt := 0
	for key, value := range *finMap{
		realKey_int, err :=  Cast89En(key)
		realKey := strconv.Itoa(realKey_int)
		if err !=  nil{
			continue
		}
		valStr := ZipBase64(value.String())
		final := fmt.Sprintf("{\"mobile\": \"%v\", \"month\": \"%v\", \"val\": \"%v\"}", realKey, month, valStr)
		wb.Write([]byte(final))
		cnt ++
		if cnt > 999 && cnt % 1000 == 0{
			wb.Flush()
		}
	}
	wb.Flush()
	wf.Close()
	wg.Done()

}


func GenFin(origin, filePath, fileType string, month string){
	bufMap := make(map[string]*strings.Builder, 1000)
	f, err := os.Open(origin)
	if err != nil{
		fmt.Println("打开文件失败:", origin)
	}
	wf, err :=  os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0644)
	if err != nil{
		fmt.Println("打开文件失败:", filePath)
	}
	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		lineStr := string(line)
		idx := strings.Index(lineStr, ",")
		if idx > 0 {
			linekey := lineStr[:idx]
			lineVal := lineStr[idx+1:]
			if _, ok := bufMap[linekey]; ok {
				bufMap[linekey].WriteString(",")
				bufMap[linekey].WriteString(lineVal)
			} else {
				x := strings.Builder{}
				x.WriteString(lineVal)
				bufMap[linekey] = &x
			}
		}
	}
	switch fileType {
	case "mysql":
		wg.Add(1)
		go GenMysql(wf, &bufMap, month)
	case "json":
		wg.Add(1)
		go GenJson(wf, &bufMap, month)
	default:
		panic(fileType + "格式不支持....")
	}
}