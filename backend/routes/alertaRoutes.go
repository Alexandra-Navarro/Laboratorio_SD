package routes

import (
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/config"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/controllers"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/services"
	"github.com/gin-gonic/gin"
)

func InitAlertaRoutes(api *gin.RouterGroup) {
	alertaService := services.NewAlertaService(config.DB)
	alertaController := controllers.NewAlertaController(alertaService)

	alertas := api.Group("/alertas")
	{
		alertas.POST("/", alertaController.Create)
		alertas.GET("/:id", alertaController.GetByID)
		alertas.GET("/", alertaController.GetAll)
		alertas.PUT("/:id", alertaController.Update)
		alertas.DELETE("/:id", alertaController.Delete)
	}
}
