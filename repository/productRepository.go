package repository

import (
	"training-go-products/database"
	"training-go-products/entity"
)

type Product struct {
	Database *database.Database
}

func (this *Product) Get(id uint) (*entity.Product, error) {
	panic("Not implemented")
}

func (this *Product) GetListOfProducts() (*entity.Product, error) {
	panic("Not implemented")
}

func (this *Product) Add() (*entity.Product, error) {
	panic("Not implemented")
}

func (this *Product) Delete(id uint) (*entity.Product, error) {
	panic("Not implemented")
}
