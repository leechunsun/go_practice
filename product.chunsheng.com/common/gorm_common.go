package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 在使用连接是需要加锁
var coreMap = make(map[string]*gorm.DB, 10)

func initDBConnection(alias, addr string) {
	conn, err := gorm.Open("mysql", addr)
	if err != nil {
		fmt.Println("数据库连接失败 >>> alias:", alias, " dburl:", addr, " err:", err)
		return
	}
	conn.DB().SetMaxIdleConns(5)
	conn.DB().SetMaxOpenConns(10)
	coreMap[alias] = conn
}

func GetConnection(alias string) *gorm.DB {
	conn := coreMap[alias]
	if conn == nil {
		return nil
	}
	conn.DB().Ping()
	return conn
}

func GetDefaultConnection() *gorm.DB {
	return GetConnection("default")
}

func init() {
	initDBConnection("default", "root:rootroot@tcp(127.0.0.1:3306)/gopro?charset=utf8&parseTime=True&loc=Local")
}
