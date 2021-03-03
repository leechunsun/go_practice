package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	hellocase "hellocase/proto"
)

type Hellocase struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Hellocase) Call(ctx context.Context, req *hellocase.Request, rsp *hellocase.Response) error {
	log.Info("Received Hellocase.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Hellocase) Stream(ctx context.Context, req *hellocase.StreamingRequest, stream hellocase.Hellocase_StreamStream) error {
	log.Infof("Received Hellocase.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&hellocase.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Hellocase) PingPong(ctx context.Context, stream hellocase.Hellocase_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&hellocase.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
