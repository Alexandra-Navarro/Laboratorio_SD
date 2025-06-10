package controllers

import (
	"net/http"
	"strconv"

	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/models"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/services"
	"github.com/gin-gonic/gin"
)

type SalaController struct {
	SalaService *services.SalaService
}

func NewSalaController(service *services.SalaService) *SalaController {
	return &SalaController{SalaService: service}
}

func (sc *SalaController) Create(c *gin.Context) {
	var sala models.Sala
	if err := c.ShouldBindJSON(&sala); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := sc.SalaService.CreateSala(&sala); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, sala)
}

func (sc *SalaController) GetByID(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	sala, err := sc.SalaService.GetSalaByID(uint(idUint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sala no encontrada"})
		return
	}
	c.JSON(http.StatusOK, sala)
}

func (sc *SalaController) GetAll(c *gin.Context) {
	salas, err := sc.SalaService.GetAllSalas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, salas)
}

func (sc *SalaController) Update(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var sala models.Sala
	if err := c.ShouldBindJSON(&sala); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	updatedSala, err := sc.SalaService.UpdateSala(uint(idUint), &sala)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sala = *updatedSala

	c.JSON(http.StatusOK, sala)
}

func (sc *SalaController) Delete(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := sc.SalaService.DeleteSala(uint(idUint)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sala no encontrada"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
