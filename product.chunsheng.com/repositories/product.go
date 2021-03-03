package repositories

import (
	"github.com/jinzhu/gorm"
	"product.chunsheng.com/common"
	"product.chunsheng.com/datamodels"
)

type ProductRepositories interface {
	Insert(*datamodels.Product) (int64, error)
	Delete(int64) bool
	Update(product *datamodels.Product) bool
	FindById(int64) *datamodels.Product
	FindAll() []datamodels.Product
}


type ProductRepositoriesImpl struct {
	db *gorm.DB
}

func NewProductRepositories() ProductRepositories{
	db := common.GetDefaultConnection()
	return &ProductRepositoriesImpl{db:db}
}

func (m *ProductRepositoriesImpl) Insert(product *datamodels.Product) (int64, error){
	res := m.db.Create(product)
	return product.ID, res.Error
}

func (m *ProductRepositoriesImpl) Delete(pid int64) bool {
	res := m.db.Delete(&datamodels.Product{}, pid)
	return res.Error != nil
}

func (m *ProductRepositoriesImpl) Update(product *datamodels.Product) bool {
	res := m.db.Model(product).Update(&product)
	return res.Error != nil
}

func (m *ProductRepositoriesImpl) FindById(pid int64) *datamodels.Product  {
	var product datamodels.Product
	m.db.Where("id = ?", pid).First(&product)
	return &product
}

func (m *ProductRepositoriesImpl) FindAll() []datamodels.Product {
	var pros []datamodels.Product
	m.db.Find(&pros)
	return pros
}
