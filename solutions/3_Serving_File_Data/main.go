package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

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

func handleRoot(c *gin.Context) {
	c.JSON(418, gin.H{"status": "I'm a teapot 🫖"})
}

func handleSquirrels(c *gin.Context) {
	fileContents, err := os.ReadFile("../../static/data.json")
	if err != nil {
		panic(err)
	}

	squirrels := []Squirrel{}
	err = json.Unmarshal(fileContents, &squirrels)
	if err != nil {
		panic(err)
	}

	c.JSON(200, squirrels)
}

func main() {
	r := gin.Default()
	r.GET("/", handleRoot)
	r.GET("/squirrels", handleSquirrels)

	r.Run("localhost:4321")
}
