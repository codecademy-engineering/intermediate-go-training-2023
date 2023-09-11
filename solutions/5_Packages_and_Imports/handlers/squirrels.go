package handlers

import (
	"github.com/codecademy-engineering/intermediate-go-training-2023/models"
	"github.com/gin-gonic/gin"
)

type SquirrelHandler struct {
	data []models.Squirrel
}

func NewSquirrelHandler(data []models.Squirrel) *SquirrelHandler {
	return &SquirrelHandler{
		data: data,
	}
}

func (h *SquirrelHandler) GetAll(c *gin.Context) {
	c.JSON(200, h.data)
}

func (h *SquirrelHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	for _, s := range h.data {
		if s.ID == id {
			c.JSON(200, s)
			return
		}
	}
	c.JSON(404, nil)
}
