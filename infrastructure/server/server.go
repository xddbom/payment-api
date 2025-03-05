package server

import (
	"crypto/tls"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

type Config struct {
	Port     string
	CertFile string
	KeyFile  string
}

func NewServer(cfg Config, handler http.Handler) *http.Server {
	server := &http.Server{
		Addr:    cfg.Port,
		Handler: handler,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	if err := http2.ConfigureServer(server, &http2.Server{}); err != nil {
		log.Fatalf("Failed to configure HTTP/2: %v", err)
	}
	return server
}
