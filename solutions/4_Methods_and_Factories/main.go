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

type SquirrelSighting struct {
	SquirrelID             string   `json:"Squirrel_ID"`
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
	data []SquirrelSighting
}

func (h *SquirrelHandler) GetAll(c *gin.Context) {
	c.JSON(200, h.data)
}

func NewSquirrelHandler(data []SquirrelSighting) *SquirrelHandler {
	return &SquirrelHandler{
		data: data,
	}
}

func loadSquirrelData() ([]SquirrelSighting, error) {

	fileContents, err := os.ReadFile("../../static/data.json")
	if err != nil {
		return nil, err
	}

	sightings := []SquirrelSighting{}
	err = json.Unmarshal(fileContents, &sightings)
	if err != nil {
		return nil, err
	}

	return sightings, nil
}

func main() {
	data, err := loadSquirrelData()
	if err != nil {
		panic(err)
	}
	squirrelHandler := NewSquirrelHandler(data)

	r := gin.Default()
	r.GET("/", handleRoot)
	r.GET("/squirrels", squirrelHandler.All)

	r.Run("localhost:4321")
}
