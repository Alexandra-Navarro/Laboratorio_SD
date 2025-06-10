package routes

import (
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/config"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/controllers"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/services"
	"github.com/gin-gonic/gin"
)

func InitUsuarioRoutes(api *gin.RouterGroup) {
	usuarioService := services.NewUsuarioService(config.DB)
	usuarioController := controllers.NewUsuarioController(usuarioService)

	usuarios := api.Group("/usuarios")
	{
		usuarios.POST("/", usuarioController.Create)
		usuarios.GET("/:rut", usuarioController.GetByID)
		usuarios.GET("/", usuarioController.GetAll)
		usuarios.PUT("/:rut", usuarioController.Update)
		usuarios.DELETE("/:rut", usuarioController.Delete)
	}
}
