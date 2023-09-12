package main

import (
	"github.com/codecademy-engineering/intermediate-go-training-2023/handlers"
	"github.com/codecademy-engineering/intermediate-go-training-2023/sos"
	"github.com/gin-gonic/gin"
)

func main() {
	mockBeacon := sos.NewMockBeacon(nil)
	sh := handlers.NewSquirrelHandler(mockBeacon)

	r := gin.Default()
	r.GET("/", handlers.HandleRoot)
	r.GET("/squirrels", sh.GetAll)
	r.GET("/squirrels/:id", sh.GetById)
	r.POST("/squirrels/:id/sos", sh.SOS)

	r.Run("localhost:4321")
}
