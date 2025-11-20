package main

import (
	"QA-service/internal/config"
	"QA-service/internal/router"
	"QA-service/internal/storage"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("Starting QA Service...")

	cfg := config.Load()

	dsn := cfg.GetDSN()
	log.Printf("Attempting to connect with DSN: host=%s, port=%s, dbname=%s, user=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBUser)

	storage, err := storage.NewStorage(dsn)
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}
	defer storage.Close()

	// Проверка подключения к БД
	if err := storage.HealthCheck(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to database!")

	r := router.NewRouter(storage)

	log.Printf("Server starting on port %s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, r))

}
