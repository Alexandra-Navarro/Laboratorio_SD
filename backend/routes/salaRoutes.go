package routes

import (
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/config"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/controllers"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/services"
	"github.com/gin-gonic/gin"
)

func InitSalaRoutes(api *gin.RouterGroup) {
	salaService := services.NewSalaService(config.DB)
	salaController := controllers.NewSalaController(salaService)

	salas := api.Group("/salas")
	{
		salas.POST("/", salaController.Create)
		salas.GET("/", salaController.GetAll)
		salas.GET("/:id", salaController.GetByID)
		salas.PUT("/:id", salaController.Update)
		salas.DELETE("/:id", salaController.Delete)
	}
}
