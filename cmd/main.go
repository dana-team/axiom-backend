package main

import (
	"github.com/dana-team/axiom-backend/internal/routes"
	"github.com/dana-team/axiom-backend/internal/utils"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load .env file")
	}

	logger := initializeLogger()
	defer syncLogger(logger)

	mongoClient, err := utils.InitMongoDB()
	if err != nil {
		logger.Fatal("Failed to initialize MongoDB client", zap.Error(err))
	}
	defer mongoClient.Disconnect()

	router := gin.Default()
	routes.SetupClusterRoutes(router, mongoClient)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Info("Starting server", zap.String("port", port))
	if err := router.Run(":" + port); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}

func initializeLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	return logger
}

func syncLogger(logger *zap.Logger) {
	if err := logger.Sync(); err != nil {
		log.Fatalf("Failed to sync logger: %v", err)
	}
}
