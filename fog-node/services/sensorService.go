package services

import (
	"database/sql"
	"fmt"

	"fog-node/models"
)

// SensorService maneja la lógica de negocio relacionada con los sensores
type SensorService struct {
	db *sql.DB
}

// NewSensorService crea una nueva instancia del servicio
func NewSensorService(db *sql.DB) *SensorService {
	return &SensorService{db: db}
}

// ProcessSensorData procesa los datos recibidos de un sensor
func (s *SensorService) ProcessSensorData(data models.SensorData, salaID int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("error al iniciar transacción: %v", err)
	}
	defer tx.Rollback()

	// Obtener IDs de variables ambientales
	var tempID, humID, co2ID, noiseID, lightID int
	if err := s.getVariableIDs(tx, &tempID, &humID, &co2ID, &noiseID, &lightID); err != nil {
		return err
	}

	// Obtener IDs de sensores
	var tempSensorID, humSensorID, co2SensorID, noiseSensorID, lightSensorID int
	if err := s.getSensorIDs(tx, salaID, tempID, humID, co2ID, noiseID, lightID,
		&tempSensorID, &humSensorID, &co2SensorID, &noiseSensorID, &lightSensorID); err != nil {
		return err
	}

	// Insertar mediciones
	if err := s.insertMeasurements(tx, data, tempSensorID, humSensorID, co2SensorID, noiseSensorID, lightSensorID); err != nil {
		return err
	}

	// Verificar umbrales
	if err := s.checkThresholds(tx, salaID, data); err != nil {
		return err
	}

	return tx.Commit()
}

// GetRecentMeasurements obtiene las mediciones recientes de una sala
func (s *SensorService) GetRecentMeasurements(salaID string) ([]models.Medicion, error) {
	rows, err := s.db.Query(`
		SELECT m.fecha, v.nombre, m.valor, v.unidad_medida
		FROM medicion m
		JOIN sensor s ON m.sensor_id = s.id
		JOIN variable_ambiental v ON s.variable_id = v.id
		WHERE s.sala_id = $1
		ORDER BY m.fecha DESC
		LIMIT 100`, salaID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mediciones []models.Medicion
	for rows.Next() {
		var m models.Medicion
		if err := rows.Scan(&m.Fecha, &m.Variable, &m.Valor, &m.UnidadMedida); err != nil {
			return nil, err
		}
		mediciones = append(mediciones, m)
	}

	return mediciones, nil
}

// GetActiveAlerts obtiene las alertas activas de una sala
func (s *SensorService) GetActiveAlerts(salaID string) ([]models.Alerta, error) {
	rows, err := s.db.Query(`
		SELECT tipo, descripcion, valor_detectado, umbral, fecha
		FROM alerta
		WHERE sala_id = $1
		ORDER BY fecha DESC
		LIMIT 50`, salaID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alertas []models.Alerta
	for rows.Next() {
		var a models.Alerta
		if err := rows.Scan(&a.Tipo, &a.Descripcion, &a.ValorDetectado, &a.Umbral, &a.Fecha); err != nil {
			return nil, err
		}
		alertas = append(alertas, a)
	}

	return alertas, nil
}

// Métodos privados auxiliares
func (s *SensorService) getVariableIDs(tx *sql.Tx, tempID, humID, co2ID, noiseID, lightID *int) error {
	var err error
	err = tx.QueryRow("SELECT id FROM variable_ambiental WHERE nombre = 'Temperatura'").Scan(tempID)
	if err != nil {
		return fmt.Errorf("error al obtener ID de variable temperatura: %v", err)
	}
	err = tx.QueryRow("SELECT id FROM variable_ambiental WHERE nombre = 'Humedad'").Scan(humID)
	if err != nil {
		return fmt.Errorf("error al obtener ID de variable humedad: %v", err)
	}
	err = tx.QueryRow("SELECT id FROM variable_ambiental WHERE nombre = 'CO2'").Scan(co2ID)
	if err != nil {
		return fmt.Errorf("error al obtener ID de variable CO2: %v", err)
	}
	err = tx.QueryRow("SELECT id FROM variable_ambiental WHERE nombre = 'Ruido'").Scan(noiseID)
	if err != nil {
		return fmt.Errorf("error al obtener ID de variable ruido: %v", err)
	}
	err = tx.QueryRow("SELECT id FROM variable_ambiental WHERE nombre = 'Iluminación'").Scan(lightID)
	if err != nil {
		return fmt.Errorf("error al obtener ID de variable iluminación: %v", err)
	}
	return nil
}

func (s *SensorService) getSensorIDs(tx *sql.Tx, salaID, tempID, humID, co2ID, noiseID, lightID int,
	tempSensorID, humSensorID, co2SensorID, noiseSensorID, lightSensorID *int) error {
	var err error
	err = tx.QueryRow("SELECT id FROM sensor WHERE sala_id = $1 AND variable_id = $2", salaID, tempID).Scan(tempSensorID)
	if err != nil {
		return fmt.Errorf("error al obtener ID de sensor de temperatura: %v", err)
	}
	err = tx.QueryRow("SELECT id FROM sensor WHERE sala_id = $1 AND variable_id = $2", salaID, humID).Scan(humSensorID)
	if err != nil {
		return fmt.Errorf("error al obtener ID de sensor de humedad: %v", err)
	}
	err = tx.QueryRow("SELECT id FROM sensor WHERE sala_id = $1 AND variable_id = $2", salaID, co2ID).Scan(co2SensorID)
	if err != nil {
		return fmt.Errorf("error al obtener ID de sensor de CO2: %v", err)
	}
	err = tx.QueryRow("SELECT id FROM sensor WHERE sala_id = $1 AND variable_id = $2", salaID, noiseID).Scan(noiseSensorID)
	if err != nil {
		return fmt.Errorf("error al obtener ID de sensor de ruido: %v", err)
	}
	err = tx.QueryRow("SELECT id FROM sensor WHERE sala_id = $1 AND variable_id = $2", salaID, lightID).Scan(lightSensorID)
	if err != nil {
		return fmt.Errorf("error al obtener ID de sensor de iluminación: %v", err)
	}
	return nil
}

func (s *SensorService) insertMeasurements(tx *sql.Tx, data models.SensorData,
	tempSensorID, humSensorID, co2SensorID, noiseSensorID, lightSensorID int) error {
	var err error
	_, err = tx.Exec("INSERT INTO medicion (fecha, valor, sensor_id) VALUES ($1, $2, $3)",
		data.Timestamp, data.Temperature, tempSensorID)
	if err != nil {
		return fmt.Errorf("error al insertar medición de temperatura: %v", err)
	}

	_, err = tx.Exec("INSERT INTO medicion (fecha, valor, sensor_id) VALUES ($1, $2, $3)",
		data.Timestamp, data.Humidity, humSensorID)
	if err != nil {
		return fmt.Errorf("error al insertar medición de humedad: %v", err)
	}

	_, err = tx.Exec("INSERT INTO medicion (fecha, valor, sensor_id) VALUES ($1, $2, $3)",
		data.Timestamp, data.CO2, co2SensorID)
	if err != nil {
		return fmt.Errorf("error al insertar medición de CO2: %v", err)
	}

	_, err = tx.Exec("INSERT INTO medicion (fecha, valor, sensor_id) VALUES ($1, $2, $3)",
		data.Timestamp, data.Noise, noiseSensorID)
	if err != nil {
		return fmt.Errorf("error al insertar medición de ruido: %v", err)
	}

	_, err = tx.Exec("INSERT INTO medicion (fecha, valor, sensor_id) VALUES ($1, $2, $3)",
		data.Timestamp, data.Light, lightSensorID)
	if err != nil {
		return fmt.Errorf("error al insertar medición de iluminación: %v", err)
	}

	return nil
}

func (s *SensorService) checkThresholds(tx *sql.Tx, salaID int, data models.SensorData) error {
	// Verificar temperatura
	var tempUmbralBajo, tempUmbralAlto float64
	err := tx.QueryRow("SELECT umbral_bajo, umbral_alto FROM variable_ambiental WHERE nombre = 'Temperatura'").Scan(&tempUmbralBajo, &tempUmbralAlto)
	if err != nil {
		return fmt.Errorf("error al obtener umbrales de temperatura: %v", err)
	}

	if float64(data.Temperature) < tempUmbralBajo {
		_, err = tx.Exec("INSERT INTO alerta (tipo, descripcion, valor_detectado, umbral, sala_id) VALUES ($1, $2, $3, $4, $5)",
			"preventivo", "Temperatura por debajo del umbral", data.Temperature, "bajo", salaID)
		if err != nil {
			return fmt.Errorf("error al crear alerta de temperatura baja: %v", err)
		}
	} else if float64(data.Temperature) > tempUmbralAlto {
		_, err = tx.Exec("INSERT INTO alerta (tipo, descripcion, valor_detectado, umbral, sala_id) VALUES ($1, $2, $3, $4, $5)",
			"preventivo", "Temperatura por encima del umbral", data.Temperature, "alto", salaID)
		if err != nil {
			return fmt.Errorf("error al crear alerta de temperatura alta: %v", err)
		}
	}

	// Verificar CO2
	var co2UmbralAlto float64
	err = tx.QueryRow("SELECT umbral_alto FROM variable_ambiental WHERE nombre = 'CO2'").Scan(&co2UmbralAlto)
	if err != nil {
		return fmt.Errorf("error al obtener umbral de CO2: %v", err)
	}

	if float64(data.CO2) > co2UmbralAlto {
		_, err = tx.Exec("INSERT INTO alerta (tipo, descripcion, valor_detectado, umbral, sala_id) VALUES ($1, $2, $3, $4, $5)",
			"critico", "Nivel de CO2 crítico", data.CO2, "alto", salaID)
		if err != nil {
			return fmt.Errorf("error al crear alerta de CO2: %v", err)
		}
	}

	return nil
}
