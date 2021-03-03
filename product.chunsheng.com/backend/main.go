package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	_ "product.chunsheng.com/common"
	"product.chunsheng.com/repositories"
)

func main() {
	app := iris.New()
	// 设置日志级别
	app.Logger().SetLevel("debug")
	// 设置模板文件位置
	template := iris.HTML("./backend/web/views", ".html").Layout("layout/layout.html").Reload(true)
	app.RegisterView(template)
	// 设置静态文件访问位置
	app.HandleDir("/assets", "./backend/web/assets")
	// 设置异常跳转
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "页面异常。。。"))
		ctx.View("layout/error.html")
	})
	// 注册视图
	product := app.Party("/product")
	productCon := mvc.Configure(product)
	productCon.Register()



	p := repositories.NewProductRepositories()
	for _, prod := range p.FindAll(){
		fmt.Println(prod)
	}
	// 启动服务
	app.Run(iris.Addr("localhost:8080"))
}
