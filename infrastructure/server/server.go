package server

import (
    "crypto/tls"
    "net/http"
    "golang.org/x/net/http2"
)

type Config struct {
    Port         string
    CertFile     string
    KeyFile      string
}

func NewServer(cfg Config, handler http.Handler) *http.Server {
    server := &http.Server{
        Addr:    cfg.Port,
        Handler: handler,
        TLSConfig: &tls.Config{
            MinVersion: tls.VersionTLS12,
        },
    }
    http2.ConfigureServer(server, &http2.Server{})
    return server
}