package services

import (
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/models"

	"gorm.io/gorm"
)

type EscuelaService struct {
	DB *gorm.DB
}

func NewEscuelaService(db *gorm.DB) *EscuelaService {
	return &EscuelaService{DB: db}
}

func (s *EscuelaService) CreateEscuela(escuela *models.Escuela) error {
	if err := s.DB.Create(&escuela).Error; err != nil {
		return err
	}
	return nil
}

func (s *EscuelaService) GetEscuelaByID(id uint) (*models.Escuela, error) {
	var escuela models.Escuela
	if err := s.DB.First(&escuela, id).Error; err != nil {
		return nil, err
	}
	return &escuela, nil
}

func (s *EscuelaService) GetAllEscuelas() ([]models.Escuela, error) {
	var escuelas []models.Escuela
	if err := s.DB.Find(&escuelas).Error; err != nil {
		return nil, err
	}
	return escuelas, nil
}

func (s *EscuelaService) UpdateEscuela(id uint, newEscuela *models.Escuela) (*models.Escuela, error) {
	var escuela models.Escuela
	if err := s.DB.First(&escuela, id).Error; err != nil {
		return nil, err
	}

	escuela.Nombre = newEscuela.Nombre
	escuela.Direccion = newEscuela.Direccion
	escuela.Comuna = newEscuela.Comuna

	if err := s.DB.Save(&escuela).Error; err != nil {
		return nil, err
	}
	return &escuela, nil
}

func (s *EscuelaService) DeleteEscuela(id uint) error {
	if err := s.DB.Delete(&models.Escuela{}, id).Error; err != nil {
		return err
	}
	return nil
}
