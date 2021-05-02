package main

import (
	"backend/conf"
	"backend/server"
)

func main() {
	// load config
	conf.Init()

	// load router
	r := server.NewRouter()
	r.Run(":9000")
}
