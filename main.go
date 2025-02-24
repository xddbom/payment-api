package main

import (
	"github.com/gin-gonic/gin"

	"payment-api/delivery/http"
)

func main() {
	r := gin.Default()

	http.RoutesSetup(r)

	r.Run(":8080")
}
