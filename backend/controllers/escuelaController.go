package controllers

import (
	"net/http"
	"strconv"

	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/models"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/services"
	"github.com/gin-gonic/gin"
)

type EscuelaController struct {
	EscuelaService *services.EscuelaService
}

func NewEscuelaController(service *services.EscuelaService) *EscuelaController {
	return &EscuelaController{EscuelaService: service}
}

func (ec *EscuelaController) Create(c *gin.Context) {
	var escuela models.Escuela
	if err := c.ShouldBindJSON(&escuela); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ec.EscuelaService.CreateEscuela(&escuela); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, escuela)
}

func (ec *EscuelaController) GetByID(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	escuela, err := ec.EscuelaService.GetEscuelaByID(uint(idUint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Escuela no encontrada"})
		return
	}

	c.JSON(http.StatusOK, escuela)
}

func (ec *EscuelaController) GetAll(c *gin.Context) {
	escuelas, err := ec.EscuelaService.GetAllEscuelas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, escuelas)
}

func (ec *EscuelaController) Update(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var escuela models.Escuela
	if err := c.ShouldBindJSON(&escuela); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEscuela, err := ec.EscuelaService.UpdateEscuela(uint(idUint), &escuela)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	escuela = *updatedEscuela

	c.JSON(http.StatusOK, escuela)
}

func (ec *EscuelaController) Delete(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := ec.EscuelaService.DeleteEscuela(uint(idUint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
