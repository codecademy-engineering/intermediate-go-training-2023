package main

import (
	"github.com/gin-gonic/gin"
)

type StatusResponse struct {
	Status string `json:"status"`
}

func handleRoot(c *gin.Context) {
	status := StatusResponse{Status: "I'm a teapot 🫖"}
	c.JSON(418, status)
}

func main() {
	r := gin.Default()
	r.GET("/", handleRoot)

	r.Run("localhost:4321")
}
