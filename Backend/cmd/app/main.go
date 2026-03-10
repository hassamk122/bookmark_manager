package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hassamk122/bookmark_manager/config"
	dbconfig "github.com/hassamk122/bookmark_manager/config/db_config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	db := dbconfig.ConnectToDB(config.DatabaseUrl)
	defer db.Close()

	serverAddr := fmt.Sprintf(":%s", config.ServerPort)

	server := &http.Server{
		Addr: serverAddr,
	}

	log.Printf("Starting server on Port: %s ...", serverAddr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed. Reason: %v", err)
	}
}
