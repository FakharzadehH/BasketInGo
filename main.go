package main

import (
	"log"

	"github.com/FakharzadehH/BasketInGo/internal/config"
	"github.com/FakharzadehH/BasketInGo/internal/logger"
	"github.com/FakharzadehH/BasketInGo/internal/server"
)

func main() {
	if err := config.Load("config.yaml"); err != nil {
		log.Fatal("error loading config")
	}
	logger.Init()
	logger.Logger().Fatal(server.Start())
}
