package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

type StatusResponse struct {
	Status string
}

type furColor struct {
	Primary    string
	Highlights string
}

type coordinates struct {
	Lat float64
	Lng float64
}

type Squirrel struct {
	ID                     string
	FurColor               furColor
	Park                   string
	InteractionsWithHumans string
	Activities             []string
	Coorditantes           coordinates
}

type SquirrelHandler struct {
	squirrels []Squirrel
}

func NewSquirrelHandler() *SquirrelHandler {
	fileContents, err := os.ReadFile("../../static/data.json")
	if err != nil {
		panic(err)
	}

	squirrels := []Squirrel{}
	err = json.Unmarshal(fileContents, &squirrels)
	if err != nil {
		panic(err)
	}

	return &SquirrelHandler{
		squirrels: squirrels,
	}
}

func (sh *SquirrelHandler) GetAll(c *gin.Context) {
	c.JSON(200, sh.squirrels)
}

func (sh *SquirrelHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	for _, s := range sh.squirrels {
		if id == s.ID {
			c.JSON(200, s)
			return
		}
	}
	c.JSON(404, nil)
}

func handleRoot(c *gin.Context) {
	status := StatusResponse{Status: "I'm a teapot 🫖"}
	c.JSON(418, status)
}

func main() {
	sh := NewSquirrelHanlder()

	r := gin.Default()
	r.GET("/", handleRoot)
	r.GET("/squirrels", sh.GetAll)
	r.GET("/squirrels/:id", sh.GetByID)

	r.Run("localhost:4321")
}
