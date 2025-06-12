package routes

import (
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/config"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/controllers"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/services"
	"github.com/gin-gonic/gin"
)

func InitUmbralRoutes(api *gin.RouterGroup) {
	umbralService := services.NewUmbralService(config.DB)
	umbralController := controllers.NewUmbralController(umbralService)

	r := api.Group("/umbrales")
	{
		r.GET("", umbralController.GetBySalaID)
		r.POST("", umbralController.Create)
		r.PUT(":id", umbralController.Update)
		r.DELETE(":id", umbralController.Delete)
	}
}
