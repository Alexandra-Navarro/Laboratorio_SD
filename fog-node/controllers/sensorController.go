package controllers

import (
	"net/http"

	"fog-node/services"

	"github.com/gin-gonic/gin"
)

type SensorController struct {
	sensorService *services.SensorService
}

func NewSensorController(sensorService *services.SensorService) *SensorController {
	return &SensorController{
		sensorService: sensorService,
	}
}

func (c *SensorController) GetRecentMeasurements(ctx *gin.Context) {
	salaID := ctx.Param("sala_id")
	mediciones, err := c.sensorService.GetRecentMeasurements(salaID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, mediciones)
}

func (c *SensorController) GetActiveAlerts(ctx *gin.Context) {
	salaID := ctx.Param("sala_id")
	alertas, err := c.sensorService.GetActiveAlerts(salaID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, alertas)
}
