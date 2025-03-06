package main

import (
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
	"payment-api/config"
	delivery "payment-api/delivery/http"
	_ "payment-api/docs"
	"payment-api/infrastructure/database"
	"payment-api/infrastructure/server"
)

func main() {
	dbConfig := config.NewDBConfig()

	db, err := database.NewPostgresConnection(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	log.Println("Connected to database successfully")

	cfg := server.Config{
		Port:     ":8080",
		CertFile: "./infrastructure/certs/cert.pem",
		KeyFile:  "./infrastructure/certs/key.pem",
	}

	r := gin.Default()
	delivery.SetupRoutes(r)

	srv := server.NewServer(cfg, r)

	log.Println("HTTPS with HTTP/2 started on port", cfg.Port)
	log.Fatal(srv.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile))
}
