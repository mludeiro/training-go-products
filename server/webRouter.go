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

	sm.Methods(http.MethodGet).Path("/api/v1/products/{id:[0-9]+}").HandlerFunc(handler.HandlerProduct)
	sm.Methods(http.MethodGet).Path("/api/v1/products").HandlerFunc(handler.HandlerProducts)
	sm.Methods(http.MethodDelete).Path("/api/v1/products/{id:[0-9]+}").HandlerFunc(handler.HandlerProduct)
	sm.Methods(http.MethodPut).Path("/api/v1/products/{id:[0-9]+}").HandlerFunc(handler.HandlerProduct)
	sm.Methods(http.MethodPost).Path("/api/v1/products").HandlerFunc(handler.HandlerProducts)

	this.router = sm
	return this.router
}
