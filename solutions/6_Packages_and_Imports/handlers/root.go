package handlers

import (
	"github.com/gin-gonic/gin"
)

type StatusResponse struct {
	Status string `json:"status"`
}

func HandleRoot(c *gin.Context) {
	status := StatusResponse{Status: "I'm a teapot 🫖"}
	c.JSON(418, status)
}
