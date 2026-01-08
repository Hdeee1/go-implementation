package main

import (
	"log"
	"net/http"

	"github.com/Hdeee1/go-implementation/config"
	"github.com/Hdeee1/go-implementation/internal/repository"
	"github.com/Hdeee1/go-implementation/internal/handlers"
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
	router.HandleFunc("/students", handlers.CreateStudentHandler(repo)).Methods("GET")
	router.HandleFunc("/students", handlers.CreateStudentHandler(repo)).Methods("POST")

	log.Printf("Server running on :%s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, router))
}