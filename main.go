package main

import (
	"backend/initialize"
	"backend/server"
)

func main() {
	// load config
	initialize.Init()

	// load router
	server.StartServer()
}
