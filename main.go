package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/victor-lima-142/simple-financial-api/database"
	"github.com/victor-lima-142/simple-financial-api/internal/handlers"
	"github.com/victor-lima-142/simple-financial-api/internal/utils"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using defaults")
	}

	user := utils.GetVarOrDefault("MONGO_USER", "guest")
	pass := utils.GetVarOrDefault("MONGO_PASS", "strongpassword")
	host := utils.GetVarOrDefault("MONGO_HOST", "127.0.0.1")
	port := utils.GetVarOrDefault("MONGO_PORT", "27017")
	dbName := utils.GetVarOrDefault("MONGO_DB", "databaseNm")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin", user, pass, host, port, dbName)

	engine := gin.Default()

	if err := database.InitMongoDB(uri, dbName); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	handlers.NewContributorHandler(engine).Initialize()
	handlers.NewCostHandler(engine).Initialize()
	handlers.NewProjectionHandler(engine).Initialize()
	handlers.NewScenarioHandler(engine).Initialize()

	if err := engine.Run(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
