package http_handler

import (
	"fmt"
	"io/ioutil"
	"localcache/cache_model"
	"localcache/db_instance"
	"net/http"
	"strconv"
	"time"
)

type StringHandler struct {
	DB *db_instance.CacheDB
}

func (s *StringHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request){
	var real_key string
	query := req.URL.Query()
	keys := query.Get("key")
	if keys != "" || req.Method == "DELETE"{
		real_key = keys
	}else{
		resp.WriteHeader(401)
		return
	}
	if req.Method == "POST"{
		var dline int64
		if dlines := query.Get("dl"); dlines != "" {
			dlint, err := strconv.Atoi(dlines)
			if err == nil{
				dline = int64(dlint)
			}
		}
		if dline <= 0{
			dline = 5 * 60
		}
		resp.WriteHeader(200)
		body, err := ioutil.ReadAll(req.Body)
		if err == nil{
			stringmodel := cache_model.StringModel{
				Dine:     dline,
				SaveTime: time.Now().Unix(),
				Data: string(body),
			}
			s.DB.Save(real_key, stringmodel, dline)
			resp.Write([]byte(`{"success": true}`))
			return
		}
		resp.Write([]byte(`{"success": false}`))
		return
	}else if req.Method == "GET" {
		dt := s.DB.Get(real_key)
		resp.WriteHeader(200)
		if dt == nil{
			resp.Write([]byte(fmt.Sprintf(`{"success": false, "data": "over deadline or non data."}`)))
			return
		}
		s, ok := dt.(cache_model.StringModel)
		if !ok {
			resp.Write([]byte(fmt.Sprintf(`{"success": false, "data": "data type error."}`)))
			return
		}

		resp.Write([]byte(fmt.Sprintf(`{"success": true, "data": "%s"}`, s.Data)))
		return
	}else if req.Method == "DELETE"{
		s.DB.Clear()
		resp.WriteHeader(200)
		return
	}else{
		resp.WriteHeader(405)
		return
	}
}
