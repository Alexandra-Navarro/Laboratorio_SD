package routes

import (
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/config"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/controllers"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/services"
	"github.com/gin-gonic/gin"
)

func InitMedicionRoutes(api *gin.RouterGroup) {
	medicionService := services.NewMedicionService(config.DB)
	medicionController := controllers.NewMedicionController(medicionService)

	mediciones := api.Group("/mediciones")
	{
		mediciones.POST("/", medicionController.Create)
		mediciones.GET("/", medicionController.GetAll)
		mediciones.GET("/:id", medicionController.GetByID)
		mediciones.PUT("/:id", medicionController.Update)
		mediciones.DELETE("/:id", medicionController.Delete)
	}
}
