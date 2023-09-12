package handlers

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/codecademy-engineering/intermediate-go-training-2023/models"
	"github.com/codecademy-engineering/intermediate-go-training-2023/sos"
	"github.com/gin-gonic/gin"
)

type SquirrelHandler struct {
	squirrels []models.Squirrel
	sosBeacon sos.Beacon
}

func NewSquirrelHandler(sosBeacon sos.Beacon) *SquirrelHandler {
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
		sosBeacon: sosBeacon,
	}
}

func (sh *SquirrelHandler) GetAll(c *gin.Context) {
	c.JSON(200, sh.squirrels)
}

func (sh *SquirrelHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	squirrel, err := sh.lookupByID(id)
	if err != nil {
		c.JSON(404, nil)
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
	return nil, errors.New("no squirrel found matching provided ID")
}

func (sh *SquirrelHandler) SOS(c *gin.Context) {
	id := c.Param("id")
	squirrel, err := sh.lookupByID(id)
	if err != nil {
		c.JSON(404, err)
		return
	}
	err = sh.sosBeacon.AlertAllAgents(squirrel)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, gin.H{"message": "Alerted all agents of an attack!"})
}
