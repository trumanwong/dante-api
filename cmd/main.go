package main

import (
	"dante-api/api/http/routers"
	"dante-api/configs"
	"fmt"
	"log"
)

func init() {
	configs.Setup()
}

func main() {
	router := routers.InitRouter()
	err := router.Run(fmt.Sprintf(":%d", configs.Config.HttpPort))
	if err != nil {
		log.Fatalln(err)
	}
}
