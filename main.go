package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/victor-lima-142/simple-financial-api/database"
	"github.com/victor-lima-142/simple-financial-api/internal/handlers"
)

func main() {
	engine := gin.Default()
	err := database.InitMongoDB("", "")
	if err != nil {
		log.Fatal(err.Error())
	}
	handlers.NewContributorHandler(engine).Initialize()
	handlers.NewCostHandler(engine).Initialize()
	handlers.NewProjectionHandler(engine).Initialize()
	handlers.NewScenarioHandler(engine).Initialize()

	engine.Run()
}
