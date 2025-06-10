package routes

import (
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/config"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/controllers"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/services"
	"github.com/gin-gonic/gin"
)

func InitVariableAmbientalRoutes(api *gin.RouterGroup) {
	vaService := services.NewVariableAmbientalService(config.DB)
	vaController := controllers.NewVariableAmbientalController(vaService)

	variables := api.Group("/variables")
	{
		variables.POST("/", vaController.Create)
		variables.GET("/:id", vaController.GetByID)
		variables.GET("/", vaController.GetAll)
		variables.PUT("/:id", vaController.Update)
		variables.DELETE("/:id", vaController.Delete)
	}
}
