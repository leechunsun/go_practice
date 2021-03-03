package main

import (
	"com.chunsheng.kiti/endpoint"
	"com.chunsheng.kiti/midware"
	"com.chunsheng.kiti/requestAndResponse"
	"com.chunsheng.kiti/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
	"os"
)

func main() {
	svc := &service.StringServiceImpl{}
	upperEp := endpoint.MakeUpperCaseEndpoint(svc)
	logger := log.NewLogfmtLogger(os.Stdout)
	upperEp = midware.LoggingMiddleWare(log.With(logger, "method", "uppercase"))(upperEp)
	uppserver := httptransport.NewServer(upperEp,
		decodeUpperCaseRequest,
		encodeResponse)

	cntserver := httptransport.NewServer(endpoint.MakeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse)

	http.Handle("/uppercase", uppserver)
	http.Handle("/count", cntserver)
	fmt.Println("start server ....")
	fmt.Println("server fatal: ", http.ListenAndServe(":8990", nil))

}


func decodeUpperCaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requestAndResponse.UpperCaseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil{
		return nil, err
	}
	return req, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req requestAndResponse.CountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil{
		return nil, err
	}
	return req, nil
}

func encodeResponse(_ context.Context, writer http.ResponseWriter, i interface{}) error  {
	return json.NewEncoder(writer).Encode(i)
}

