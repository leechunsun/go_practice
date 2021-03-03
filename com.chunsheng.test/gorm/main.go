package main

import (
	"com.chunsheng.test/common"
	"fmt"
	"sync"
	"time"
)

type Product struct {
	ID          int64  `json:"id" sql:"ID" imooc:"id"`
	ProductName string `json:"ProductName" sql:"productName" imooc:"ProductName" gorm:"column:productName"`
	ProductNum  int64  `json:"ProductNum" sql:"productNum" imooc:"ProductNum" gorm:"column:productNum"`
	ProductImage  string `json:"ProductImage" sql:"productImage" imooc:"ProductImage" gorm:"column:productImage"`
	ProductUrl  string `json:"ProductUrl" sql:"productUrl" imooc:"ProductUrl" gorm:"column:productUrl"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Product) TableName() string {
	return "product"
}



func main() {
	wg := sync.WaitGroup{}

	for i:=0; i< 10;i++{
		wg.Add(1)
		go func(ii int) {
			db := common.GetDefaultConnection()
			if ! db.HasTable(&Product{}) {
				db.CreateTable(&Product{})
			}
			// 增加记录
			//pro := Product{
			//	ProductImage:"qwert",
			//	ProductNum: int64(ii),
			//	ProductName:"zzz",
			//	ProductUrl:"kkkk",
			//}
			//db.Create(&pro)
			//fmt.Println(pro.ID, pro.CreatedAt)

			// 查询记录
			//users := make([]Product, 1000)
			//res := db.Find(&users)
			//fmt.Println(res.RowsAffected)
			p := &Product{}
			r := db.Where("id = ?", ii + 5).First(p)
			//
			fmt.Println(p)
			fmt.Println(r)

			wg.Done()
		}(i)
	}
	wg.Wait()
}
