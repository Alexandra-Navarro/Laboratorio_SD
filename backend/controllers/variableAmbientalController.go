package controllers

import (
	"net/http"
	"strconv"

	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/models"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/services"
	"github.com/gin-gonic/gin"
)

type VariableAmbientalController struct {
	VariableAmbientalService *services.VariableAmbientalService
}

func NewVariableAmbientalController(service *services.VariableAmbientalService) *VariableAmbientalController {
	return &VariableAmbientalController{VariableAmbientalService: service}
}

func (vc *VariableAmbientalController) Create(c *gin.Context) {
	var variable models.VariableAmbiental
	if err := c.ShouldBindJSON(&variable); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := vc.VariableAmbientalService.CreateVariableAmbiental(&variable); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, variable)
}

func (vc *VariableAmbientalController) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	variable, err := vc.VariableAmbientalService.GetVariableAmbientalByID(uint(idUint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Variable ambiental no encontrada"})
		return
	}

	c.JSON(http.StatusOK, variable)
}

func (vc *VariableAmbientalController) GetAll(c *gin.Context) {
	variables, err := vc.VariableAmbientalService.GetAllVariablesAmbientales()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, variables)
}

func (vc *VariableAmbientalController) Update(c *gin.Context) {
	idStr := c.Param("id")
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var nueva models.VariableAmbiental
	if err := c.ShouldBindJSON(&nueva); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Obtener la actual
	variable, err := vc.VariableAmbientalService.GetVariableAmbientalByID(uint(idUint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Variable ambiental no encontrada"})
		return
	}

	// Actualizar campos
	variable.Nombre = nueva.Nombre
	variable.UnidadMedida = nueva.UnidadMedida
	variable.UmbralBajo = nueva.UmbralBajo
	variable.UmbralAlto = nueva.UmbralAlto

	if err := vc.VariableAmbientalService.DB.Save(&variable).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, variable)
}

func (vc *VariableAmbientalController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := vc.VariableAmbientalService.DeleteVariableAmbiental(uint(idUint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
