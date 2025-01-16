package main

import (
	"apiref/core"
	"apiref/src/util"
)

func main() {
	util.DatabaseMigrate()
	server := core.Server{Port: 8000}
	router := util.ControllerBuilder{}.Build()
	err := server.StartServer(router)
	if err != nil {
		return
	}
}
