package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

type StatusResponse struct {
	Status string `json:"status"`
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
	ID                     string   `json:"ID"`
	FurColor               furColor `json:"Fur_Color"`
	Park                   string
	InteractionsWithHumans string `json:"Interactions_With_Humans"`
	Activities             []string
	Coorditantes           coordinates
}

func handleRoot(c *gin.Context) {
	status := StatusResponse{Status: "I'm a teapot 🫖"}
	c.JSON(418, status)
}

type SquirrelHandler struct {
	data []Squirrel
}

func (h *SquirrelHandler) GetAll(c *gin.Context) {
	c.JSON(200, h.data)
}

func (h *SquirrelHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	for _, s := range h.data {
		if s.ID == id {
			c.JSON(200, s)
			return
		}
	}
	c.JSON(404, nil)
}

func NewSquirrelHandler(data []Squirrel) *SquirrelHandler {
	return &SquirrelHandler{
		data: data,
	}
}

func loadSquirrelData() ([]Squirrel, error) {

	fileContents, err := os.ReadFile("../../static/data.json")
	if err != nil {
		return nil, err
	}

	squirrels := []Squirrel{}
	err = json.Unmarshal(fileContents, &squirrels)
	if err != nil {
		return nil, err
	}

	return squirrels, nil
}

func main() {
	data, err := loadSquirrelData()
	if err != nil {
		panic(err)
	}
	squirrelHandler := NewSquirrelHandler(data)

	r := gin.Default()
	r.GET("/", handleRoot)
	r.GET("/squirrels", squirrelHandler.GetAll)
	r.GET("/squirrels/:id", squirrelHandler.GetByID)

	r.Run("localhost:4321")
}
