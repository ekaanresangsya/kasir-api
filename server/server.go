package server

import (
	"fmt"
	"log"
)

func Start() {
	log.Print("Starting server ...")

	router := InitRouter()

	port := "8080"
	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Printf("error running server, got %v", err)
	}
}
