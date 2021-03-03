package endpoint

import (
	"com.chunsheng.kiti/requestAndResponse"
	"com.chunsheng.kiti/service"
	"context"
	"github.com/go-kit/kit/endpoint"
)

func MakeUpperCaseEndpoint(srv service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
			req := request.(requestAndResponse.UpperCaseRequest)
			r, err := srv.UpperCase(req.S)
			resp := requestAndResponse.UpperCaseResponse{
				V: r,
			}
			if err != nil{
				resp.E = err.Error()
			}
			return resp, nil
	}
}

func MakeCountEndpoint(srv service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(requestAndResponse.CountRequest)
		resp := srv.Count(req.S)
		return requestAndResponse.CountResponse{V:resp}, nil
	}
}

