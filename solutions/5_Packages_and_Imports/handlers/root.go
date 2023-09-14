package handlers

import (
	"github.com/gin-gonic/gin"
)

func HandleRoot(c *gin.Context) {
	c.JSON(418, gin.H{"status": "I'm a teapot 🫖"})
}
