package services

import (
	"fmt"
	"time"

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

func (s *AlertaService) VerificarUmbrales(medicion *models.Medicion) error {
	// Obtener el sensor asociado a la medición
	var sensor models.Sensor
	if err := s.DB.First(&sensor, "id = ?", medicion.SensorID).Error; err != nil {
		return fmt.Errorf("error al obtener sensor: %w", err)
	}

	// Obtener la variable ambiental del sensor
	var variable models.VariableAmbiental
	if err := s.DB.First(&variable, "id = ?", sensor.VariableID).Error; err != nil {
		return fmt.Errorf("error al obtener variable ambiental: %w", err)
	}

	// Buscar umbral personalizado
	var umbralPersonalizado models.Umbral
	hasCustom := false
	if err := s.DB.Where("sala_id = ? AND variable_id = ?", sensor.SalaID, sensor.VariableID).First(&umbralPersonalizado).Error; err == nil {
		hasCustom = true
	}

	// Determinar umbrales a usar
	var umbralBajo, umbralAlto float64
	if hasCustom {
		umbralBajo = umbralPersonalizado.UmbralBajo
		umbralAlto = umbralPersonalizado.UmbralAlto
	} else {
		// No hay umbral definido para esta variable en esta sala
		fmt.Printf("[ADVERTENCIA] No hay umbral definido para sala_id=%d, variable_id=%d\n", sensor.SalaID, sensor.VariableID)
		return nil // O puedes retornar un error si prefieres
	}

	// Determinar el tipo de alerta según los umbrales definidos
	var tipoAlerta, descripcion, umbral string

	switch {
	case medicion.Valor < umbralBajo:
		tipoAlerta = "critico"
		descripcion = fmt.Sprintf("%s bajo el umbral mínimo", variable.Nombre)
		umbral = "bajo"
	case medicion.Valor > umbralAlto:
		tipoAlerta = "critico"
		descripcion = fmt.Sprintf("%s sobre el umbral máximo", variable.Nombre)
		umbral = "alto"
	case variable.Nombre == "Temperatura" && medicion.Valor >= 25 && medicion.Valor <= umbralAlto:
		tipoAlerta = "preventivo"
		descripcion = "Temperatura en rango preventivo"
		umbral = "variable"
	case variable.Nombre == "CO2" && medicion.Valor > 1000 && medicion.Valor <= 1500:
		tipoAlerta = "preventivo"
		descripcion = "CO₂ en rango preventivo"
		umbral = "variable"
	default:
		tipoAlerta = "informativo"
		descripcion = fmt.Sprintf("%s en condiciones normales", variable.Nombre)
		umbral = "variable"
	}

	// Crear la alerta si aplica
	if tipoAlerta != "" {
		alerta := models.Alerta{
			Tipo:           tipoAlerta,
			Descripcion:    descripcion,
			ValorDetectado: medicion.Valor,
			Umbral:         umbral,
			Fecha:          time.Now(),
			SalaID:         sensor.SalaID,
		}
		if err := s.CreateAlerta(&alerta); err != nil {
			return fmt.Errorf("error al crear alerta: %w", err)
		}
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
