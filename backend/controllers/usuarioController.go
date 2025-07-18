package controllers

import (
	"net/http"

	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/models"
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/services"
	"github.com/gin-gonic/gin"
)

type UsuarioController struct {
	UsuarioService *services.UsuarioService
}

func NewUsuarioController(service *services.UsuarioService) *UsuarioController {
	return &UsuarioController{UsuarioService: service}
}

func (uc *UsuarioController) Create(c *gin.Context) {
	var usuario models.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.UsuarioService.CreateUsuario(&usuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, usuario)
}

func (uc *UsuarioController) GetByID(c *gin.Context) {
	rut := c.Param("rut")

	var usuario models.Usuario
	if err := uc.UsuarioService.DB.First(&usuario, "rut_usuario = ?", rut).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func (uc *UsuarioController) GetAll(c *gin.Context) {
	usuarios, err := uc.UsuarioService.GetAllUsuarios()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, usuarios)
}

func (uc *UsuarioController) Update(c *gin.Context) {
	rut := c.Param("rut")

	var nuevoUsuario models.Usuario
	if err := c.ShouldBindJSON(&nuevoUsuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var usuario models.Usuario
	if err := uc.UsuarioService.DB.First(&usuario, "rut_usuario = ?", rut).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	usuario.Nombre = nuevoUsuario.Nombre
	usuario.Email = nuevoUsuario.Email
	usuario.Password = nuevoUsuario.Password
	usuario.Rol = nuevoUsuario.Rol
	usuario.EscuelaID = nuevoUsuario.EscuelaID

	if err := uc.UsuarioService.DB.Save(&usuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func (uc *UsuarioController) Delete(c *gin.Context) {
	rut := c.Param("rut")

	if err := uc.UsuarioService.DB.Delete(&models.Usuario{}, "rut_usuario = ?", rut).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (uc *UsuarioController) Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de login inválidos"})
		return
	}

	usuario, err := uc.UsuarioService.GetByEmail(loginData.Email)
	if err != nil || usuario == nil || usuario.Password != loginData.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	// token, err := config.GenerateJWT(usuario.Rut)
	// if err != nil {
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el token"})
	//     return
	// }

	// c.JSON(http.StatusOK, gin.H{"token": token})

	c.JSON(http.StatusOK, gin.H{
		"message": "Login exitoso",
		"usuario": usuario,
	})
}
