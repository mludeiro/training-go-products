package server

import (
	"net/http"

	"training-go-products/handler"

	"github.com/gorilla/mux"
)

type WebRouter struct {
	router *mux.Router
}

func (this *WebRouter) GetRouter() *mux.Router {
	if this.router != nil {
		return this.router
	}

	sm := mux.NewRouter()

	sm.Methods(http.MethodGet).Path("/products/{id:[0-9]+}").HandlerFunc(handler.HandlerProduct)
	sm.Methods(http.MethodGet).Path("/products").HandlerFunc(handler.HandlerProducts)
	sm.Methods(http.MethodDelete).Path("/products/{id:[0-9]+}").HandlerFunc(handler.HandlerProduct)
	sm.Methods(http.MethodPut).Path("/products/{id:[0-9]+}").HandlerFunc(handler.HandlerProduct)
	sm.Methods(http.MethodPost).Path("/products").HandlerFunc(handler.HandlerProducts)

	this.router = sm
	return this.router
}
