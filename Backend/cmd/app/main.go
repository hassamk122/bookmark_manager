package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hassamk122/bookmark_manager/config"
	dbconfig "github.com/hassamk122/bookmark_manager/config/db_config"
	"github.com/hassamk122/bookmark_manager/internals/handlers"
	"github.com/hassamk122/bookmark_manager/internals/repos"
	"github.com/hassamk122/bookmark_manager/internals/routes"
	"github.com/hassamk122/bookmark_manager/internals/services"
	"github.com/hassamk122/bookmark_manager/internals/store"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	db := dbconfig.ConnectToDB(config.DatabaseUrl)
	defer db.Close()

	queries := store.New(db)

	userRepo := repos.NewUserRepo(queries)
	userService := services.NewUserService(db, userRepo)

	serverAddr := fmt.Sprintf(":%s", config.ServerPort)

	handler := handlers.NewHandler(userService)

	mux := http.NewServeMux()
	routes.SetupRoutes(mux, handler)

	server := &http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}

	log.Printf("Starting server on Port: %s ...", serverAddr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed. Reason: %v", err)
	}
}
