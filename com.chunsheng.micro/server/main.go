package main

import (
	_ "github.com/kataras/iris/v12/mvc"
	"github.com/micro/micro/v3/service"
	_ "github.com/micro/micro/v3/service/registry/mdns"
	_ "github.com/kataras/iris/v12"
)

func main() {

	serv := service.New(
		service.Name("chunsheng"),
		service.Version("v1.0"),
		)

	serv.Server()


}


