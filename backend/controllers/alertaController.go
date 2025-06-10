package controllers

import (
	"net/http"
	"strconv"

	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/models"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/services"
	"github.com/gin-gonic/gin"
)

type AlertaController struct {
	AlertaService *services.AlertaService
}

func NewAlertaController(service *services.AlertaService) *AlertaController {
	return &AlertaController{AlertaService: service}
}

func (ac *AlertaController) Create(c *gin.Context) {
	var alerta models.Alerta
	if err := c.ShouldBindJSON(&alerta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ac.AlertaService.CreateAlerta(&alerta); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, alerta)
}

func (ac *AlertaController) GetByID(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	alerta, err := ac.AlertaService.GetAlertaByID(uint(idUint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Alerta no encontrada"})
		return
	}

	c.JSON(http.StatusOK, alerta)
}

func (ac *AlertaController) GetAll(c *gin.Context) {
	alertas, err := ac.AlertaService.GetAllAlertas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, alertas)
}

func (ac *AlertaController) Update(c *gin.Context) {
	id := c.Param("id")
	var alerta models.Alerta
	if err := c.ShouldBindJSON(&alerta); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	updatedAlerta, err := ac.AlertaService.UpdateAlerta(uint(idUint), &alerta)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	alerta = *updatedAlerta

	c.JSON(http.StatusOK, alerta)
}

func (ac *AlertaController) Delete(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := ac.AlertaService.DeleteAlerta(uint(idUint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
