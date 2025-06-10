package routes

import (
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/config"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/controllers"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/services"
	"github.com/gin-gonic/gin"
)

func InitEscuelaRoutes(api *gin.RouterGroup) {
	escuelaService := services.NewEscuelaService(config.DB)
	escuelaController := controllers.NewEscuelaController(escuelaService)

	escuelas := api.Group("/escuelas")
	{
		escuelas.POST("/", escuelaController.Create)
		escuelas.GET("/", escuelaController.GetAll)
		escuelas.GET("/:id", escuelaController.GetByID)
		escuelas.PUT("/:id", escuelaController.Update)
		escuelas.DELETE("/:id", escuelaController.Delete)
	}
}
