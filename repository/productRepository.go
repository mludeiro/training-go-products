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

func InsertProduct(product entity.Product) (uint, error) {
	panic("Not implemented")
}

func GetProduct(productID uint) (*entity.Product, error) {
	panic("Not implemented")
}

func UpdateProduct(product entity.Product) error {
	result := db.GetDB().Model(&product).Update("name", product.Name)
	return result.Error
}

func RemoveProduct(productID uint) error {
	panic("Not implemented")
}
