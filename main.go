package main

import (
	"log"

	"github.com/gin-gonic/gin"

	delivery "payment-api/delivery/http"
	_ "payment-api/docs"
	"payment-api/infrastructure/server"
)

func main() {
	r := gin.Default()
	delivery.SetupRoutes(r)

	cfg := server.Config{
		Port:     ":8080",
		CertFile: "./infrastructure/certs/cert.pem",
		KeyFile:  "./infrastructure/certs/key.pem",
	}

	srv := server.NewServer(cfg, r)

	log.Println("HTTPS with HTTP/2 started on port", cfg.Port)
	log.Fatal(srv.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile))
}
