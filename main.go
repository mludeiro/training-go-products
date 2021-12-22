package main

import (
	"fmt"
	"net/http"
	"training-go-products/handler"
)

const productsPath = "products"

func main() {
	productsHandler := http.HandlerFunc(handler.HandlerProducts)
	productHandler := http.HandlerFunc(handler.HandlerProduct)
	http.Handle(fmt.Sprintf("%s", productsPath), productsHandler)
	http.Handle(fmt.Sprintf("%s/", productsPath), productHandler)
}
