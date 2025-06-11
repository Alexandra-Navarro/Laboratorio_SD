package routes

import (
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/config"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/controllers"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/services"
	"github.com/gin-gonic/gin"
)

func InitMedicionRoutes(api *gin.RouterGroup) {
	alertaService := services.NewAlertaService(config.DB)
	medicionService := services.NewMedicionService(config.DB, alertaService)
	medicionController := controllers.NewMedicionController(medicionService)

	mediciones := api.Group("/mediciones")
	{
		mediciones.POST("/", medicionController.CreateMedicion)
		mediciones.GET("/sala/:sala_id", medicionController.GetMedicionesBySala)
		mediciones.GET("/variable/:variable_id", medicionController.GetMedicionesByVariable)
	}
}
