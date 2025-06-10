package routes

import (
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/config"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/controllers"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/services"
	"github.com/gin-gonic/gin"
)

func InitSensorRoutes(api *gin.RouterGroup) {
	sensorService := services.NewSensorService(config.DB)
	sensorController := controllers.NewSensorController(sensorService)

	sensores := api.Group("/sensores")
	{
		sensores.POST("/", sensorController.Create)
		sensores.GET("/:id", sensorController.GetByID)
		sensores.GET("/", sensorController.GetAll)
		sensores.PUT("/:id", sensorController.Update)
		sensores.DELETE("/:id", sensorController.Delete)
	}
}
