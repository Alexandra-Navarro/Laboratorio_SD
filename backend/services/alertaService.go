package services

import (
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/models"

	"gorm.io/gorm"
)

type AlertaService struct {
	DB *gorm.DB
}

func NewAlertaService(db *gorm.DB) *AlertaService {
	return &AlertaService{DB: db}
}

func (s *AlertaService) CreateAlerta(alerta *models.Alerta) error {
	if err := s.DB.Create(&alerta).Error; err != nil {
		return err
	}
	return nil
}

func (s *AlertaService) GetAlertaByID(id uint) (*models.Alerta, error) {
	var alerta models.Alerta
	if err := s.DB.First(&alerta, id).Error; err != nil {
		return nil, err
	}
	return &alerta, nil
}

func (s *AlertaService) GetAllAlertas() ([]models.Alerta, error) {
	var alertas []models.Alerta
	if err := s.DB.Find(&alertas).Error; err != nil {
		return nil, err
	}
	return alertas, nil
}

func (s *AlertaService) UpdateAlerta(id uint, newAlerta *models.Alerta) (*models.Alerta, error) {
	var alerta models.Alerta
	if err := s.DB.First(&alerta, id).Error; err != nil {
		return nil, err
	}

	alerta.Tipo = newAlerta.Tipo
	alerta.Descripcion = newAlerta.Descripcion
	alerta.ValorDetectado = newAlerta.ValorDetectado
	alerta.Umbral = newAlerta.Umbral
	alerta.Fecha = newAlerta.Fecha
	alerta.SalaID = newAlerta.SalaID

	if err := s.DB.Save(&alerta).Error; err != nil {
		return nil, err
	}
	return &alerta, nil
}

func (s *AlertaService) DeleteAlerta(id uint) error {
	if err := s.DB.Delete(&models.Alerta{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *AlertaService) GetAlertasBySalaID(salaID uint) ([]models.Alerta, error) {
	var alertas []models.Alerta
	if err := s.DB.Where("sala_id = ?", salaID).Find(&alertas).Error; err != nil {
		return nil, err
	}
	return alertas, nil
}
