package main

import (
	"hellocase/handler"
	pb "hellocase/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("hellocase"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterHellocaseHandler(srv.Server(), new(handler.Hellocase))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
