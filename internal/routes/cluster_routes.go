package routes

import (
	"github.com/dana-team/axiom-backend/internal/controllers"
	"github.com/dana-team/axiom-backend/internal/middleware"
	"github.com/dana-team/axiom-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

// SetupClusterRoutes configures all the HTTP routes for cluster-related endpoints,
// sets up middleware for CORS, logging, and error handling, and initializes the
// cluster controller with the provided MongoDB client
func SetupClusterRoutes(router *gin.Engine, mongoClient *utils.MongoClient) {
	router.Use(middleware.CORS())
	router.Use(middleware.Logger())
	router.Use(middleware.ErrorHandler())

	clusterController := controllers.NewClusterController(mongoClient)
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	v1 := router.Group("v1")
	{
		clusters := v1.Group("/clusters")
		{
			clusters.GET("", clusterController.GetClusters)
		}
	}
}
