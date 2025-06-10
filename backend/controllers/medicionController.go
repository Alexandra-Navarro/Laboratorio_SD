package controllers

import (
	"net/http"
	"strconv"

	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/models"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/services"
	"github.com/gin-gonic/gin"
)

type MedicionController struct {
	MedicionService *services.MedicionService
}

func NewMedicionController(service *services.MedicionService) *MedicionController {
	return &MedicionController{MedicionService: service}
}

func (mc *MedicionController) Create(c *gin.Context) {
	var medicion models.Medicion
	if err := c.ShouldBindJSON(&medicion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := mc.MedicionService.CreateMedicion(&medicion); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, medicion)
}

func (mc *MedicionController) GetByID(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	medicion, err := mc.MedicionService.GetMedicionByID(uint(idUint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medición no encontrada"})
		return
	}

	c.JSON(http.StatusOK, medicion)
}

func (mc *MedicionController) GetAll(c *gin.Context) {
	mediciones, err := mc.MedicionService.GetAllMediciones()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mediciones)
}

func (mc *MedicionController) Update(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var medicion models.Medicion
	if err := c.ShouldBindJSON(&medicion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedMedicion, err := mc.MedicionService.UpdateMedicion(uint(idUint), &medicion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	medicion = *updatedMedicion

	c.JSON(http.StatusOK, medicion)
}

func (mc *MedicionController) Delete(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := mc.MedicionService.DeleteMedicion(uint(idUint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
