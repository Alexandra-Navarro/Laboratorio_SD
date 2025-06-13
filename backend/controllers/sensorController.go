package controllers

import (
	"net/http"
	"strconv"

	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/models"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/services"
	"github.com/gin-gonic/gin"
)

type SensorController struct {
	SensorService *services.SensorService
}

func NewSensorController(service *services.SensorService) *SensorController {
	return &SensorController{SensorService: service}
}

func (sc *SensorController) Create(c *gin.Context) {
	var sensor models.Sensor
	if err := c.ShouldBindJSON(&sensor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := sc.SensorService.CreateSensor(&sensor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, sensor)
}

func (sc *SensorController) GetByID(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	sensor, err := sc.SensorService.GetSensorByID(uint(idUint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sensor no encontrado"})
		return
	}

	c.JSON(http.StatusOK, sensor)
}

func (sc *SensorController) GetAll(c *gin.Context) {
	sensores, err := sc.SensorService.GetAllSensors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sensores)
}

func (sc *SensorController) Update(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var sensor models.Sensor
	if err := c.ShouldBindJSON(&sensor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedSensor, err := sc.SensorService.UpdateSensor(uint(idUint), &sensor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedSensor)
}

func (sc *SensorController) Delete(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := sc.SensorService.DeleteSensor(uint(idUint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (sc *SensorController) GetSensorsBySala(c *gin.Context) {
	salaID := c.Param("sala_id")
	salaIDInt, err := strconv.Atoi(salaID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de sala inválido"})
		return
	}

	sensors, err := sc.SensorService.GetSensorsBySala(salaIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sensors)
}
