package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
