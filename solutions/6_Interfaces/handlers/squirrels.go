package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/codecademy-engineering/intermediate-go-training-2023/alert"
	"github.com/codecademy-engineering/intermediate-go-training-2023/models"
	"github.com/gin-gonic/gin"
)

type SquirrelHandler struct {
	squirrels []models.Squirrel
	beacon    alert.Beacon
}

func NewSquirrelHandler(beacon alert.Beacon) *SquirrelHandler {
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
		beacon:    beacon,
	}
}

func (sh *SquirrelHandler) GetAll(c *gin.Context) {
	c.JSON(200, sh.squirrels)
}

func (sh *SquirrelHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	squirrel, err := sh.lookupByID(id)
	if err != nil {
		c.JSON(404, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, squirrel)
}

func (sh *SquirrelHandler) lookupByID(id string) (*models.Squirrel, error) {
	for _, s := range sh.squirrels {
		if id == s.ID {
			return &s, nil
		}
	}
	return nil, errors.New("squirrel not found")
}

func (sh *SquirrelHandler) ThreatAlert(c *gin.Context) {
	id := c.Param("id")
	squirrel, err := sh.lookupByID(id)
	if err != nil {
		c.JSON(404, gin.H{"message": err.Error()})
		return
	}
	err = sh.beacon.AlertAllAgents(squirrel)
	if err != nil {
		c.JSON(500, gin.H{"message": fmt.Sprintf("beacon failure: %w", err)})
		return
	}
	c.JSON(200, gin.H{"message": fmt.Sprintf("Alerted all agents that squirrel %s is an imminent threat!", squirrel.ID)})
}
