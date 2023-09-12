package handlers

import (
	"encoding/json"
	"os"

	"github.com/codecademy-engineering/intermediate-go-training-2023/models"
	"github.com/gin-gonic/gin"
)

type SquirrelHandler struct {
	squirrels []models.Squirrel
}

func NewSquirrelHandler() *SquirrelHandler {
	fileContents, err := os.ReadFile("../../static/data.json")
	if err != nil {
		panic(err)
	}

	squirrels := []models.Squirrel{}
	err = json.Unmarshal(fileContents, &squirrels)
	if err != nil {
		panic(err)
	}

	return &SquirrelHandler{
		squirrels: squirrels,
	}
}

func (h *SquirrelHandler) GetAll(c *gin.Context) {
	c.JSON(200, h.squirrels)
}

func (h *SquirrelHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	for _, s := range h.squirrels {
		if s.ID == id {
			c.JSON(200, s)
			return
		}
	}
	c.JSON(404, nil)
}
