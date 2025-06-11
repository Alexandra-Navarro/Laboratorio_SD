package services

import (
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

// VerificarUmbrales verifica si una medición supera los umbrales y crea alertas si es necesario
func (s *AlertaService) VerificarUmbrales(medicion *models.Medicion) error {
	// Obtener la variable ambiental asociada a la medición
	var variable models.VariableAmbiental
	if err := s.DB.First(&variable, "id = ?", medicion.VariableID).Error; err != nil {
		return err
	}

	// Determinar el tipo de alerta según la variable y el valor
	var tipoAlerta string
	var descripcion string

	switch variable.Nombre {
	case "Temperatura":
		if medicion.Valor < 20 || medicion.Valor > 27 {
			tipoAlerta = "critico"
			descripcion = "Temperatura fuera del rango crítico"
		} else if medicion.Valor >= 25 && medicion.Valor <= 27 {
			tipoAlerta = "preventivo"
			descripcion = "Temperatura en rango preventivo"
		} else if medicion.Valor >= 20 && medicion.Valor <= 24 {
			tipoAlerta = "informativo"
			descripcion = "Temperatura en rango informativo"
		}

	case "Humedad":
		if medicion.Valor < 30 || medicion.Valor > 70 {
			tipoAlerta = "critico"
			descripcion = "Humedad fuera del rango crítico"
		} else if (medicion.Valor >= 30 && medicion.Valor < 40) || (medicion.Valor > 60 && medicion.Valor <= 70) {
			tipoAlerta = "preventivo"
			descripcion = "Humedad en rango preventivo"
		} else if medicion.Valor >= 40 && medicion.Valor <= 60 {
			tipoAlerta = "informativo"
			descripcion = "Humedad en rango informativo"
		}

	case "CO2":
		if medicion.Valor >= 1500 {
			tipoAlerta = "critico"
			descripcion = "Nivel de CO2 crítico"
		} else if medicion.Valor > 1000 && medicion.Valor < 1500 {
			tipoAlerta = "preventivo"
			descripcion = "Nivel de CO2 en rango preventivo"
		} else if medicion.Valor >= 400 && medicion.Valor <= 1000 {
			tipoAlerta = "informativo"
			descripcion = "Nivel de CO2 en rango informativo"
		}

	case "Ruido":
		if medicion.Valor > 55 {
			tipoAlerta = "critico"
			descripcion = "Nivel de ruido crítico"
		} else if medicion.Valor >= 36 && medicion.Valor <= 55 {
			tipoAlerta = "preventivo"
			descripcion = "Nivel de ruido en rango preventivo"
		} else if medicion.Valor < 35 {
			tipoAlerta = "informativo"
			descripcion = "Nivel de ruido en rango informativo"
		}
	}

	// Si se detectó una alerta, crearla
	if tipoAlerta != "" {
		alerta := models.Alerta{
			Tipo:           tipoAlerta,
			Descripcion:    descripcion,
			ValorDetectado: medicion.Valor,
			Umbral:         "variable",
			Fecha:          time.Now(),
			SalaID:         medicion.SalaID,
		}
		if err := s.CreateAlerta(&alerta); err != nil {
			return err
		}
	}

	return nil
}
