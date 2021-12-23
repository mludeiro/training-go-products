package server

import (
	"net/http"
	"training-go-products/tools"
)

type WebServer struct {
	server http.Server
	Router WebRouter
}

func (this *WebServer) CreateServer() {
	this.server = http.Server{
		Addr:    ":5002",                 // configure the bind address
		Handler: this.Router.GetRouter(), // set the default handler
	}

	tools.GetLogger().Println("Starting server on port 5002")

	err := this.server.ListenAndServe()
	if err != nil {
		tools.GetLogger().Printf("Error starting server: %s\n", err)
	}
}
