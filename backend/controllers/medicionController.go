package controllers

import (
	"net/http"
	"strconv"

	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/models"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/services"
	"github.com/gin-gonic/gin"
)

type MedicionController struct {
	medicionService *services.MedicionService
}

func NewMedicionController(medicionService *services.MedicionService) *MedicionController {
	return &MedicionController{
		medicionService: medicionService,
	}
}

// CreateMedicion maneja la creación de una nueva medición
func (c *MedicionController) CreateMedicion(ctx *gin.Context) {
	var medicion models.Medicion
	if err := ctx.ShouldBindJSON(&medicion); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.medicionService.CreateMedicion(&medicion); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, medicion)
}

// GetMedicionesBySala obtiene las mediciones de una sala específica
func (c *MedicionController) GetMedicionesBySala(ctx *gin.Context) {
	salaID, err := strconv.Atoi(ctx.Param("sala_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de sala inválido"})
		return
	}

	mediciones, err := c.medicionService.GetMedicionesBySala(salaID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, mediciones)
}

// GetMedicionesByVariable obtiene las mediciones de una variable específica
func (c *MedicionController) GetMedicionesByVariable(ctx *gin.Context) {
	variableID, err := strconv.Atoi(ctx.Param("variable_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de variable inválido"})
		return
	}

	mediciones, err := c.medicionService.GetMedicionesByVariable(variableID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, mediciones)
}

func (mc *MedicionController) GetByID(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	medicion, err := mc.medicionService.GetMedicionByID(uint(idUint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medición no encontrada"})
		return
	}

	c.JSON(http.StatusOK, medicion)
}

func (mc *MedicionController) GetAll(c *gin.Context) {
	mediciones, err := mc.medicionService.GetAllMediciones()
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

	updatedMedicion, err := mc.medicionService.UpdateMedicion(uint(idUint), &medicion)
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

	if err := mc.medicionService.DeleteMedicion(uint(idUint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
