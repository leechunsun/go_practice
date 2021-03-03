package datamodels

import "time"

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
