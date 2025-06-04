package services

import (
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/models"

	"gorm.io/gorm"
)

type MedicionService struct {
	DB *gorm.DB
}

func NewMedicionService(db *gorm.DB) *MedicionService {
	return &MedicionService{DB: db}
}

func (s *MedicionService) CreateMedicion(medicion *models.Medicion) error {
	if err := s.DB.Create(&medicion).Error; err != nil {
		return err
	}
	return nil
}

func (s *MedicionService) GetMedicionByID(id uint) (*models.Medicion, error) {
	var medicion models.Medicion
	if err := s.DB.First(&medicion, id).Error; err != nil {
		return nil, err
	}
	return &medicion, nil
}

func (s *MedicionService) GetAllMediciones() ([]models.Medicion, error) {
	var mediciones []models.Medicion
	if err := s.DB.Find(&mediciones).Error; err != nil {
		return nil, err
	}
	return mediciones, nil
}

func (s *MedicionService) UpdateMedicion(id uint, newMedicion *models.Medicion) (*models.Medicion, error) {
	var medicion models.Medicion
	if err := s.DB.First(&medicion, id).Error; err != nil {
		return nil, err
	}

	medicion.Fecha = newMedicion.Fecha
	medicion.Valor = newMedicion.Valor
	medicion.SensorID = newMedicion.SensorID

	if err := s.DB.Save(&medicion).Error; err != nil {
		return nil, err
	}
	return &medicion, nil
}

func (s *MedicionService) DeleteMedicion(id uint) error {
	if err := s.DB.Delete(&models.Medicion{}, id).Error; err != nil {
		return err
	}
	return nil
}
