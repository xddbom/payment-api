package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"

	delivery "payment-api/delivery/http"
	_ "payment-api/docs"
)

func main() {
	r := gin.Default()
	delivery.SetupRoutes(r)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	go func() {
		log.Println("HTTP started on port :8081")
		if err := r.Run(":8081"); err != nil {
			log.Printf("HTTP stoped. %v", err)
		}
	}()

	http2.ConfigureServer(server, &http2.Server{})
	log.Println("HTTPS with HTTP/v2 started on port :8080")
	log.Fatal(server.ListenAndServeTLS("./infrastructure/certs/cert.pem", "./infrastructure/certs/key.pem"))
}
