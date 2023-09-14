package main

import (
	"github.com/gin-gonic/gin"
)

func handleRoot(c *gin.Context) {
	c.JSON(418, gin.H{"status": "I'm a teapot 🫖"})
}

func main() {
	r := gin.Default()
	r.GET("/", handleRoot)

	r.Run("localhost:4321")
}
