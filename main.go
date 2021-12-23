package main

import (
	"training-go-products/container"
)

func main() {
	cont := container.NewContainer()

	cont.DataBase.InitializeSqlite().Migrate().CreateTestData()
	cont.WebServer.CreateServer()
}
