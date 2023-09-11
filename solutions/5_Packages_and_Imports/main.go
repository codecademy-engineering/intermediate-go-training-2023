package main

import (
	"github.com/codecademy-engineering/intermediate-go-training-2023/handlers"
	"github.com/codecademy-engineering/intermediate-go-training-2023/loaders"
	"github.com/gin-gonic/gin"
)

func main() {
	data, err := loaders.LoadSquirrelData()
	if err != nil {
		panic(err)
	}
	squirrelHandler := handlers.NewSquirrelHandler(data)

	r := gin.Default()
	r.GET("/", handlers.HandleRoot)
	r.GET("/squirrels", squirrelHandler.GetAll)
	r.GET("/squirrels/:id", squirrelHandler.GetById)

	r.Run("localhost:4321")
}
