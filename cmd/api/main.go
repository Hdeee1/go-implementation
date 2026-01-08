package main

import (
	"log"
	"net/http"

	"github.com/Hdeee1/go-implementation/config"
	"github.com/Hdeee1/go-implementation/internal/handlers"
	"github.com/Hdeee1/go-implementation/internal/middleware"
	"github.com/Hdeee1/go-implementation/internal/repository"
	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	db, err := repository.InitDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect Database:", err)
	}
	defer db.Close()

	repo := repository.NewStudentRepo(db)

	router := mux.NewRouter()
	router.Use(middleware.LoggerMiddleware)
	middleware.AuthMiddleware(handlers.CreateStudentHandler(repo))
	router.HandleFunc("/students", handlers.CreateStudentHandler(repo)).Methods("POST")

	log.Printf("Server running on :%s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, router))
}