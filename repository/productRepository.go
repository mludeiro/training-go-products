package repository

import (
	"training-go-products/database"
	"training-go-products/entity"
)

var db database.Database = *database.GetInstance()

func GetProductList() ([]*entity.Product, error) {
	var products []*entity.Product
	result := db.GetDB().Find(&products)
	return products, result.Error
}

func InsertProduct(product entity.Product) (*entity.Product, error) {
	result := db.GetDB().Create(&product)
	return &product, result.Error
}

func GetProduct(productID uint) (*entity.Product, error) {
	var product *entity.Product
	result := db.GetDB().First(&product, productID)
	return product, result.Error
}

func UpdateProduct(product entity.Product) error {
	result := db.GetDB().Model(&product).Update("name", product.Name)
	return result.Error
}

func RemoveProduct(productID uint) error {
	//missing logic for invoices
	result := db.GetDB().Delete(&entity.Product{}, productID)
	return result.Error
}
