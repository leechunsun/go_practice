package sysintit

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
)

func dbInit(alias ...string){
	if len(alias) > 0{
		for _, ali := range alias{
			doInit(ali)
		}
	}else{
		doInit("w")
	}
}

func doInit(ali string) {
	dbAlias := ali
	if ali == "w" || ali == "default" || ali == "" {
		dbAlias = "default"
		ali = "w"
	}

	host := beego.AppConfig.String(fmt.Sprintf("db_%s_host", ali))
	user := beego.AppConfig.String(fmt.Sprintf("db_%s_user", ali))
	password := beego.AppConfig.String(fmt.Sprintf("db_%s_password", ali))
	port := beego.AppConfig.String(fmt.Sprintf("db_%s_port", ali))
	database := beego.AppConfig.String(fmt.Sprintf("db_%s_database", ali))
	charset := beego.AppConfig.String(fmt.Sprintf("db_%s_charset", ali))

	dbString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", user, password, host, port, database, charset)
	err := orm.RegisterDataBase(dbAlias, "mysql", dbString, 30)
	if dbAlias == "default"{
		mod := beego.AppConfig.String("runmode") == "dev"
		orm.RunSyncdb("mysql", false, mod)
	}
	if err != nil{
		fmt.Println("初始化数据库：", dbAlias, " Err:", err)
	}

}

