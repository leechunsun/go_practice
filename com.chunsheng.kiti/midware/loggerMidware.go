package midware

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

func LoggingMiddleWare(logger log.Logger) endpoint.Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			logger.Log("calling endpoint...")
			defer logger.Log("callend endpoint...")
			response, err = e(ctx, request)
			logger.Log("cal rs: ", response, "  e: ", err)
			return
		}
	}
}
