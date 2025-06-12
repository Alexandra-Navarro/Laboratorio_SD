package services

import (
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/models"
	"gorm.io/gorm"
)

type UsuarioService struct {
	DB *gorm.DB
}

func NewUsuarioService(db *gorm.DB) *UsuarioService {
	return &UsuarioService{DB: db}
}

func (s *UsuarioService) CreateUsuario(usuario *models.Usuario) error {
	if err := s.DB.Create(&usuario).Error; err != nil {
		return err
	}
	return nil
}

func (s *UsuarioService) GetUsuarioByID(id uint) (*models.Usuario, error) {
	var usuario models.Usuario
	if err := s.DB.First(&usuario, id).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (s *UsuarioService) GetAllUsuarios() ([]models.Usuario, error) {
	var usuarios []models.Usuario
	if err := s.DB.Find(&usuarios).Error; err != nil {
		return nil, err
	}
	return usuarios, nil
}

func (s *UsuarioService) UpdateUsuario(id uint, newUsuario *models.Usuario) (*models.Usuario, error) {
	var usuario models.Usuario
	if err := s.DB.First(&usuario, id).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (s *UsuarioService) DeleteUsuario(id uint) error {
	if err := s.DB.Delete(&models.Usuario{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *UsuarioService) GetByUsername(username string) (*models.Usuario, error) {
	var usuario models.Usuario
	if err := s.DB.Where("username = ?", username).First(&usuario).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (s *UsuarioService) GetByEmail(email string) (*models.Usuario, error) {
	var usuario models.Usuario
	if err := s.DB.Where("email = ?", email).First(&usuario).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}
