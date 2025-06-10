package routes

import (
	"github.com/gin-gonic/gin"
)

// InitAllRoutes registra todas las rutas del sistema en un solo lugar.
func InitAllRoutes(r *gin.Engine) {
	api := r.Group("/api")

	InitAlertaRoutes(api)
	InitSensorRoutes(api)
	InitUsuarioRoutes(api)
	InitVariableAmbientalRoutes(api)
	InitEscuelaRoutes(api)
	InitMedicionRoutes(api)
	InitSalaRoutes(api)
}
