package main

import (
	"log"

	"github.com/msegeya56/eCommerce-smart_cads/pkg/config"
	"github.com/msegeya56/eCommerce-smart_cads/pkg/di"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Load configuration failed")
	}

	server, err := di.InitiateAPI(config)

	if err != nil {
		log.Fatal("Failed to start server")
	}
	server.Run()

}
