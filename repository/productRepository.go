package repository

import (
	"training-go-products/database"
	"training-go-products/entity"
)

type Product struct {
	Database *database.Database
}

func GetProductList() ([]*entity.Product, error) {
	panic("Not implemented")
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
	panic("Not implemented")
}
