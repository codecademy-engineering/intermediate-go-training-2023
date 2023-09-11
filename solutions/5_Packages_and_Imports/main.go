package main

import (
	"github.com/codecademy-engineering/intermediate-go-training-2023/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	sh := handlers.NewSquirrelHandler()

	r := gin.Default()
	r.GET("/", handlers.HandleRoot)
	r.GET("/squirrels", sh.GetAll)
	r.GET("/squirrels/:id", sh.GetById)

	r.Run("localhost:4321")
}
