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
	panic("Not implemented")
}

func UpdateProduct(product entity.Product) error {
	panic("Not implemented")
}

func RemoveProduct(productID uint) error {
	panic("Not implemented")
}
