package services

import (
	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/models"
	"gorm.io/gorm"
)

type SensorService struct {
	DB *gorm.DB
}

func NewSensorService(db *gorm.DB) *SensorService {
	return &SensorService{DB: db}
}

func (s *SensorService) CreateSensor(sensor *models.Sensor) error {
	if err := s.DB.Create(&sensor).Error; err != nil {
		return err
	}
	return nil
}

func (s *SensorService) GetSensorByID(id uint) (*models.Sensor, error) {
	var sensor models.Sensor
	if err := s.DB.First(&sensor, id).Error; err != nil {
		return nil, err
	}
	return &sensor, nil
}

func (s *SensorService) GetAllSensors() ([]models.Sensor, error) {
	var sensors []models.Sensor
	if err := s.DB.Find(&sensors).Error; err != nil {
		return nil, err
	}
	return sensors, nil
}

func (s *SensorService) UpdateSensor(id uint, newSensor *models.Sensor) (*models.Sensor, error) {
	var sensor models.Sensor
	if err := s.DB.First(&sensor, id).Error; err != nil {
		return nil, err
	}
	return &sensor, nil
}

func (s *SensorService) DeleteSensor(id uint) error {
	if err := s.DB.Delete(&models.Sensor{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *SensorService) GetSensorsBySala(salaID int) ([]models.Sensor, error) {
	var sensors []models.Sensor
	if err := s.DB.Where("sala_id = ?", salaID).Find(&sensors).Error; err != nil {
		return nil, err
	}
	return sensors, nil
}
