package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

type MyFile struct {
	FileName string
	PreFix string
	Content string
	Month string
	FileType string
}

func GenJsonStr(preFix, content, month string) string {
	return fmt.Sprintf("{\"mobile\": \"%v\", \"month\": \"%v\", \"val\": \"%v\"}\n", preFix, month, content)
}

func GenMysqlStr(preFix, content, month string) string {
	return preFix + "," + month + "," + content + "\n"
}


func Write(){
	fl := sync.Map{}
	fc := sync.Map{}
	for myFile := range WChan{
		go func(mf *MyFile) {
			enPrei, err := Cast89En(mf.PreFix)
			if err != nil{
				return
			}
			enPre := strconv.Itoa(enPrei)
			enCon := ZipBase64(mf.Content)
			fn := myFile.FileName
			if flc, _ := fl.LoadOrStore(fn, make(chan bool, 1));true{
				flc.(chan bool) <- true
				if sbb, _ := fc.LoadOrStore(fn, &strings.Builder{}); true{
					switch myFile.FileType {
					case "mysql":
						sbb.(*strings.Builder).WriteString(GenMysqlStr(enPre, enCon, myFile.Month))
					case "json":
						sbb.(*strings.Builder).WriteString(GenJsonStr(enPre, enCon, myFile.Month))
					default:
						return
					}
					if sbb.(*strings.Builder).Len() > 1000{
						f, _ := os.OpenFile(myFile.FileName, os.O_APPEND | os.O_CREATE| os.O_WRONLY, 0644)
						fb := bufio.NewWriter(f)
						fb.WriteString(sbb.(*strings.Builder).String())
						fb.Flush()
						f.Close()
						sbb.(*strings.Builder).Reset()
					}
				}
				<- flc.(chan bool)
			}
		}(myFile)
	}
	fc.Range(func(key, value interface{}) bool {
		f, _ := os.OpenFile(key.(string), os.O_APPEND | os.O_CREATE| os.O_WRONLY, 0644)
		fb := bufio.NewWriter(f)
		fb.WriteString(value.(*strings.Builder).String())
		fb.Flush()
		f.Close()
		return true
	})
}


func Gen(origin, finalFile, fileType , month string){
	cmd := exec.Command("/bin/bash", "-c", "sort -o "+origin+" "+ origin)
	fmt.Println("sortfile:", origin, "......")
	err := cmd.Run()
	if err != nil{
		return
	}
	f, err := os.Open(origin)
	defer f.Close()
	reader := bufio.NewReader(f)
	currentPre := ""
	sb := strings.Builder{}
	for {
		line, _, err := reader.ReadLine()
		if err != nil{
			break
		}
		lineStr := string(line)
		idx := strings.Index(lineStr, ",")
		prefix := lineStr[:idx]
		content := lineStr[idx+1:]
		if prefix != currentPre{
			if currentPre != ""{
				WChan <- &MyFile{finalFile, currentPre, sb.String(), month, fileType}
			}
			currentPre = prefix
			sb.Reset()
		}
		sb.WriteString(content)
		sb.WriteString(",")
	}
}
