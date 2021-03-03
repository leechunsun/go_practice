package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"strconv"
)

func Cast89En(origin string)(int, error) {
	// 将origin转成8进制
	i, e := strconv.Atoi(origin)
	if e != nil{
		return i, e
	}
	res := 0
	base := 1
	for i != 0 {
		v := i % 8
		i = i / 8
		res += base * v
		base *= 9
	}
	return res, nil
}

func ZipBase64(origin string) string {
	var in bytes.Buffer
	zn, _ := zlib.NewWriterLevel(&in, zlib.BestSpeed)
	zn.Write([]byte(origin))
	zn.Close()
	return base64.StdEncoding.EncodeToString(in.Bytes())
}
