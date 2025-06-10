package services

import (
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/models"
	"gorm.io/gorm"
)

type VariableAmbientalService struct {
	DB *gorm.DB
}

func NewVariableAmbientalService(db *gorm.DB) *VariableAmbientalService {
	return &VariableAmbientalService{DB: db}
}

func (s *VariableAmbientalService) CreateVariableAmbiental(variableAmbiental *models.VariableAmbiental) error {
	if err := s.DB.Create(&variableAmbiental).Error; err != nil {
		return err
	}
	return nil
}

func (s *VariableAmbientalService) GetVariableAmbientalByID(id uint) (*models.VariableAmbiental, error) {
	var variableAmbiental models.VariableAmbiental
	if err := s.DB.First(&variableAmbiental, id).Error; err != nil {
		return nil, err
	}
	return &variableAmbiental, nil
}

func (s *VariableAmbientalService) GetAllVariablesAmbientales() ([]models.VariableAmbiental, error) {
	var variablesAmbientales []models.VariableAmbiental
	if err := s.DB.Find(&variablesAmbientales).Error; err != nil {
		return nil, err
	}
	return variablesAmbientales, nil
}

func (s *VariableAmbientalService) UpdateVariableAmbiental(id uint, newVariableAmbiental *models.VariableAmbiental) (*models.VariableAmbiental, error) {
	var variableAmbiental models.VariableAmbiental
	if err := s.DB.First(&variableAmbiental, id).Error; err != nil {
		return nil, err
	}
	return &variableAmbiental, nil
}

func (s *VariableAmbientalService) DeleteVariableAmbiental(id uint) error {
	if err := s.DB.Delete(&models.VariableAmbiental{}, id).Error; err != nil {
		return err
	}
	return nil
}
