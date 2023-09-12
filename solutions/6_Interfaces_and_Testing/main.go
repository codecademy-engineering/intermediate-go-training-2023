package main

import (
	"github.com/codecademy-engineering/intermediate-go-training-2023/alert"
	"github.com/codecademy-engineering/intermediate-go-training-2023/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	mockBeacon := alert.NewMockBeacon(nil)
	sh := handlers.NewSquirrelHandler(mockBeacon)

	r := gin.Default()
	r.GET("/", handlers.HandleRoot)
	r.GET("/squirrels", sh.GetAll)
	r.GET("/squirrels/:id", sh.GetById)
	r.POST("/squirrel_alert/:id", sh.Alert)

	r.Run("localhost:4321")
}
