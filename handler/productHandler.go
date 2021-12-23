package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"training-go-products/entity"
	"training-go-products/repository"
)

var (
	Products []*entity.Product
	ID       int
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	productList, err := repository.GetProductList()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	j, err := json.Marshal(productList)
	if err != nil {
		log.Print(err)
	}
	_, err = w.Write(j)
	if err != nil {
		log.Print(err)
	}
}

func PostProducts(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	productID, err := repository.InsertProduct(product)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`{"productId":%d}`, productID)))
}

func GetProduct(w http.ResponseWriter, r *http.Request, productID uint) {
	product, err := repository.GetProduct(productID)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	j, err := json.Marshal(product)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = w.Write(j)
	if err != nil {
		log.Print(err)
	}
}

func PutProduct(w http.ResponseWriter, r *http.Request, productID uint) {
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if product.ID != productID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = repository.UpdateProduct(product)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request, productID uint) {
	err := repository.RemoveProduct(productID)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
