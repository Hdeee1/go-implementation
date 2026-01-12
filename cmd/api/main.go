package main

import (
	"log"
	"net/http"

	"github.com/Hdeee1/go-implementation/config"
	"github.com/Hdeee1/go-implementation/internal/handlers"
	"github.com/Hdeee1/go-implementation/internal/middleware"
	"github.com/Hdeee1/go-implementation/internal/repository"
	"github.com/Hdeee1/go-implementation/migrations"
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

	if err := migrations.Migrate(db); err != nil {
		log.Fatal("Migration failed:", err)
	}

	repo := repository.NewStudentRepo(db)
	

	router := mux.NewRouter()
	router.Use(middleware.LoggerMiddleware)
	router.HandleFunc("/students", handlers.GetStudentHandler(repo)).Methods("GET")
	router.Handle("/students", middleware.AuthMiddleware(handlers.CreateStudentHandler(repo))).Methods("POST")
	router.HandleFunc("/students/export", handlers.ExportStudentHandler(repo)).Methods("GET")
	router.HandleFunc("/students/import", handlers.ImportStudentHandler(repo)).Methods("POST")

	log.Printf("Server running on :%s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, router))
}