package main

import (
	"chunsheng.isir.com/web/controllers"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)


func main() {
	app := iris.New()
	// 设置日志级别
	app.Logger().SetLevel("debug")
	// 设置模板文件位置
	app.RegisterView(iris.HTML("./web/views", "html"))
	// 注册路由
	app.Get("/{mid:string}", func(context context.Context) {
		con := controllers.MovieController{context}
		con.Get()
	})

	err := app.Run(iris.Addr("localhost:8090"))
	if err != nil{
		fmt.Println("服务已停止......")
	}

}
