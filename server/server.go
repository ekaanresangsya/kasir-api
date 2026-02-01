package server

import (
	"crud-categories/internal/database"
	"fmt"
	"log"
)

func Start() {
	log.Print("Starting server ...")

	config := LoadConfig()

	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatalf("error connecting to database, got %v", err)
	}
	defer db.Close()

	router := InitRouter(db)

	err = router.Run(fmt.Sprintf(":%s", config.ServerPort))
	if err != nil {
		log.Printf("error running server, got %v", err)
	}
}
