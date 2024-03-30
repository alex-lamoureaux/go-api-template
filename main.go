package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Print("Hello World!")

	router := gin.Default()
	router.GET("/users", getUsers)
	router.Run("localhost:8080")
}
