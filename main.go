package main

import (
	"training-go-products/container"
	"training-go-products/database"
)

var Cont container.Container

func main() {
	Cont := container.NewContainer()
	Cont.DataBase = database.GetInstance()
	Cont.WebServer.CreateServer()
}
