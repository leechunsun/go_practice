package main

import (
	"net/http"
)


type MyHandler struct {

}

func (mh *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request){
	cookie, err := req.Cookie("uid")
	if err != nil{
		w.Write([]byte("错误请求"))
		return
	}
	aa := cookie.Value
	w.Write([]byte(aa))
}

func main() {
	http.Handle("/check", &MyHandler{})
	http.ListenAndServe(":9090", nil)
}
