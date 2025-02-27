package main

import (
	"github.com/gin-gonic/gin"

	"payment-api/delivery/http"
	_ "payment-api/docs"
)

func main() {
	r := gin.Default()
	http.SetupRoutes(r)

	

	r.Run(":8080")
}
