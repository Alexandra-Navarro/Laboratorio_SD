package services

import (
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/models"
	"gorm.io/gorm"
)

type SalaService struct {
	DB *gorm.DB
}

func NewSalaService(db *gorm.DB) *SalaService {
	return &SalaService{DB: db}
}

func (s *SalaService) CreateSala(sala *models.Sala) error {
	if err := s.DB.Create(&sala).Error; err != nil {
		return err
	}
	return nil
}

func (s *SalaService) GetSalaByID(id uint) (*models.Sala, error) {
	var sala models.Sala
	if err := s.DB.First(&sala, id).Error; err != nil {
		return nil, err
	}
	return &sala, nil
}

func (s *SalaService) GetAllSalas() ([]models.Sala, error) {
	var salas []models.Sala
	if err := s.DB.Find(&salas).Error; err != nil {
		return nil, err
	}
	return salas, nil
}

func (s *SalaService) UpdateSala(id uint, newSala *models.Sala) (*models.Sala, error) {
	var sala models.Sala
	if err := s.DB.First(&sala, id).Error; err != nil {
		return nil, err
	}
	return &sala, nil
}

func (s *SalaService) DeleteSala(id uint) error {
	if err := s.DB.Delete(&models.Sala{}, id).Error; err != nil {
		return err
	}
	return nil
}
