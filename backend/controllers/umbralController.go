package controllers

import (
	"net/http"
	"strconv"

	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/models"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/services"
	"github.com/gin-gonic/gin"
)

type UmbralController struct {
	UmbralService *services.UmbralService
}

func NewUmbralController(service *services.UmbralService) *UmbralController {
	return &UmbralController{UmbralService: service}
}

// GET /api/umbrales?sala_id=X
func (uc *UmbralController) GetBySalaID(c *gin.Context) {
	salaIDStr := c.Query("sala_id")
	salaID, err := strconv.Atoi(salaIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sala_id inválido"})
		return
	}
	umbrales, err := uc.UmbralService.GetUmbralesBySalaID(salaID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, umbrales)
}

// POST /api/umbrales
func (uc *UmbralController) Create(c *gin.Context) {
	var umbral models.Umbral
	if err := c.ShouldBindJSON(&umbral); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uc.UmbralService.CreateUmbral(&umbral); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, umbral)
}

// PUT /api/umbrales/:id
func (uc *UmbralController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var umbral models.Umbral
	if err := c.ShouldBindJSON(&umbral); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uc.UmbralService.UpdateUmbral(id, &umbral); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Umbral actualizado"})
}

// DELETE /api/umbrales/:id
func (uc *UmbralController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if err := uc.UmbralService.DeleteUmbral(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
