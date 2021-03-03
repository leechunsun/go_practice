package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	o := orm.NewOrm()
	o.QueryTable("user")
	ob, err := orm.NewQueryBuilder("mysql")


	c.TplName = "index.tpl"
}
