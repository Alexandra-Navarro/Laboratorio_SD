package services

import (
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/models"
	"gorm.io/gorm"
)

type UmbralService struct {
	DB *gorm.DB
}

func NewUmbralService(db *gorm.DB) *UmbralService {
	return &UmbralService{DB: db}
}

func (s *UmbralService) GetUmbralesBySalaID(salaID int) ([]models.Umbral, error) {
	var umbrales []models.Umbral
	if err := s.DB.Where("sala_id = ?", salaID).Find(&umbrales).Error; err != nil {
		return nil, err
	}
	return umbrales, nil
}

func (s *UmbralService) CreateUmbral(umbral *models.Umbral) error {
	return s.DB.Create(umbral).Error
}

func (s *UmbralService) UpdateUmbral(id int, umbral *models.Umbral) error {
	return s.DB.Model(&models.Umbral{}).Where("id = ?", id).Updates(umbral).Error
}

func (s *UmbralService) DeleteUmbral(id int) error {
	return s.DB.Delete(&models.Umbral{}, id).Error
}

func (s *UmbralService) GetUmbralBySalaAndVariable(salaID int, variableID int) (*models.Umbral, error) {
	var umbral models.Umbral
	err := s.DB.Where("sala_id = ? AND variable_id = ?", salaID, variableID).First(&umbral).Error
	if err != nil {
		return nil, err
	}
	return &umbral, nil
}
