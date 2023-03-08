package main

import (
	"fmt"
	"log"

	"enerBit-system/cmd/providers"
	"enerBit-system/config"
	"enerBit-system/internal/infra/api/router"
	"github.com/labstack/echo/v4"
)

var (
	serverHost = config.Environments().ServerHost
	serverPort = config.Environments().ServerPort
)

// main
//
//	@title			Ener Bit System
//	@version		1.0.0
//	@description	Register and monitor the energy meters that have been installed in our clients' properties
//	@license.name	Andres Puello
//	@BasePath		/api
//	@schemes		http
func main() {
	container := providers.BuildContainer()

	err := container.Invoke(func(router *router.Router, server *echo.Echo) {
		router.Init()
		server.Logger.Fatal(server.Start(fmt.Sprintf("%s:%d", serverHost, serverPort)))
	})

	if err != nil {
		log.Panic(err)
	}
}
