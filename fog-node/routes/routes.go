package routes

import (
	"fog-node/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, sensorController *controllers.SensorController) {
	api := router.Group("/api")
	{
		api.GET("/mediciones/:sala_id", sensorController.GetRecentMeasurements)
		api.GET("/alertas/:sala_id", sensorController.GetActiveAlerts)
	}
}
