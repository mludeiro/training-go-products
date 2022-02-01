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
	panic("Not implemented")
}

func RemoveProduct(productID uint) error {
	//missing logic for invoices
	result := db.GetDB().Delete(&entity.Product{}, productID)
	return result.Error
}
