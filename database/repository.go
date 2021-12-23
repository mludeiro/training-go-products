package database

import "training-go-products/entity"

func GetProductList() ([]*entity.Product, error) {
	return nil, nil
}

func InsertProduct(product entity.Product) (uint, error) {
	return 0, nil
}

func GetProduct(productID uint) (*entity.Product, error) {
	var res *entity.Product
	return res, nil
}

func UpdateProduct(product entity.Product) error {
	return nil
}

func RemoveProduct(productID uint) error {
	return nil
}
