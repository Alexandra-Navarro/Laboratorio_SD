package services

import (
	"database/sql"
	"fmt"

	"fog-node/models"
)

type SensorService struct {
	db *sql.DB
}

func NewSensorService(db *sql.DB) *SensorService {
	return &SensorService{db: db}
}

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
	var tempVarID, humVarID, co2VarID, noiseVarID, lightVarID int
	err := tx.QueryRow("SELECT id FROM variable_ambiental WHERE nombre = 'Temperatura'").Scan(&tempVarID)
	if err != nil {
		return fmt.Errorf("error al obtener id de variable Temperatura: %v", err)
	}
	err = tx.QueryRow("SELECT id FROM variable_ambiental WHERE nombre = 'Humedad'").Scan(&humVarID)
	if err != nil {
		return fmt.Errorf("error al obtener id de variable Humedad: %v", err)
	}
	err = tx.QueryRow("SELECT id FROM variable_ambiental WHERE nombre = 'CO2'").Scan(&co2VarID)
	if err != nil {
		return fmt.Errorf("error al obtener id de variable CO2: %v", err)
	}
	err = tx.QueryRow("SELECT id FROM variable_ambiental WHERE nombre = 'Ruido'").Scan(&noiseVarID)
	if err != nil {
		return fmt.Errorf("error al obtener id de variable Ruido: %v", err)
	}
	err = tx.QueryRow("SELECT id FROM variable_ambiental WHERE nombre = 'Iluminación'").Scan(&lightVarID)
	if err != nil {
		return fmt.Errorf("error al obtener id de variable Iluminación: %v", err)
	}

	// Temperatura
	var tempUmbralBajo, tempUmbralAlto, tempUmbralBajoPrev, tempUmbralAltoPrev float64
	err = tx.QueryRow("SELECT umbral_bajo, umbral_alto, umbral_bajo_prev, umbral_alto_prev FROM umbrales WHERE sala_id = $1 AND variable_id = $2", salaID, tempVarID).Scan(&tempUmbralBajo, &tempUmbralAlto, &tempUmbralBajoPrev, &tempUmbralAltoPrev)
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
	} else if float64(data.Temperature) >= tempUmbralBajoPrev && float64(data.Temperature) <= tempUmbralAltoPrev {
		_, err = tx.Exec("INSERT INTO alerta (tipo, descripcion, valor_detectado, umbral, sala_id) VALUES ($1, $2, $3, $4, $5)",
			"preventivo", "Temperatura por encima del umbral", data.Temperature, "alto", salaID)
		if err != nil {
			return fmt.Errorf("error al crear alerta de temperatura alta: %v", err)
		}
	}

	// Humedad
	var humUmbralBajo, humUmbralAlto, humUmbralBajoPrev, humUmbralAltoPrev float64
	err = tx.QueryRow("SELECT umbral_bajo, umbral_alto, umbral_bajo_prev, umbral_alto_prev FROM umbrales WHERE sala_id = $1 AND variable_id = $2", salaID, humVarID).Scan(&humUmbralBajo, &humUmbralAlto, &humUmbralBajoPrev, &humUmbralAltoPrev)
	if err != nil {
		return fmt.Errorf("error al obtener umbrales de humedad: %v", err)
	}
	if float64(data.Humidity) < humUmbralBajo {
		_, err = tx.Exec("INSERT INTO alerta (tipo, descripcion, valor_detectado, umbral, sala_id) VALUES ($1, $2, $3, $4, $5)",
			"preventivo", "Humedad por debajo del umbral", data.Humidity, "bajo", salaID)
		if err != nil {
			return fmt.Errorf("error al crear alerta de humedad baja: %v", err)
		}
	} else if float64(data.Humidity) > humUmbralAlto {
		_, err = tx.Exec("INSERT INTO alerta (tipo, descripcion, valor_detectado, umbral, sala_id) VALUES ($1, $2, $3, $4, $5)",
			"preventivo", "Humedad por encima del umbral", data.Humidity, "alto", salaID)
		if err != nil {
			return fmt.Errorf("error al crear alerta de humedad alta: %v", err)
		}
	} else if (float64(data.Humidity) <= humUmbralBajoPrev && float64(data.Humidity) >= humUmbralBajo) || (float64(data.Humidity) >= humUmbralAltoPrev && float64(data.Humidity) <= humUmbralAlto) {
		_, err = tx.Exec("INSERT INTO alerta (tipo, descripcion, valor_detectado, umbral, sala_id) VALUES ($1, $2, $3, $4, $5)",
			"preventivo", "Humedad transgrede el umbral", data.Humidity, "alto", salaID)
		if err != nil {
			return fmt.Errorf("error al crear alerta de humedad alta: %v", err)
		}
	}

	// CO2
	var co2UmbralBajo, co2UmbralAlto, co2UmbralBajoPrev, co2UmbralAltoPrev float64
	err = tx.QueryRow("SELECT umbral_bajo, umbral_alto, umbral_bajo_prev, umbral_alto_prev FROM umbrales WHERE sala_id = $1 AND variable_id = $2", salaID, co2VarID).Scan(&co2UmbralBajo, &co2UmbralAlto, &co2UmbralBajoPrev, &co2UmbralAltoPrev)
	if err != nil {
		return fmt.Errorf("error al obtener umbrales de CO2: %v", err)
	}
	if float64(data.CO2) < co2UmbralBajo {
		_, err = tx.Exec("INSERT INTO alerta (tipo, descripcion, valor_detectado, umbral, sala_id) VALUES ($1, $2, $3, $4, $5)",
			"critico", "CO2 por debajo del umbral", data.CO2, "bajo", salaID)
		if err != nil {
			return fmt.Errorf("error al crear alerta de CO2 bajo: %v", err)
		}
	} else if float64(data.CO2) > co2UmbralAlto {
		_, err = tx.Exec("INSERT INTO alerta (tipo, descripcion, valor_detectado, umbral, sala_id) VALUES ($1, $2, $3, $4, $5)",
			"critico", "Nivel de CO2 crítico", data.CO2, "alto", salaID)
		if err != nil {
			return fmt.Errorf("error al crear alerta de CO2: %v", err)
		}
	} else if float64(data.CO2) >= co2UmbralBajoPrev && float64(data.CO2) <= co2UmbralAltoPrev {
		_, err = tx.Exec("INSERT INTO alerta (tipo, descripcion, valor_detectado, umbral, sala_id) VALUES ($1, $2, $3, $4, $5)",
			"preventivo", "CO2 por encima del umbral", data.CO2, "alto", salaID)
		if err != nil {
			return fmt.Errorf("error al crear alerta de CO2 alta: %v", err)
		}
	}

	// Ruido
	var noiseUmbralBajo, noiseUmbralAlto, noiseUmbralBajoPrev, noiseUmbralAltoPrev float64
	err = tx.QueryRow("SELECT umbral_bajo, umbral_alto, umbral_bajo_prev, umbral_alto_prev FROM umbrales WHERE sala_id = $1 AND variable_id = $2", salaID, noiseVarID).Scan(&noiseUmbralBajo, &noiseUmbralAlto, &noiseUmbralBajoPrev, &noiseUmbralAltoPrev)
	if err != nil {
		return fmt.Errorf("error al obtener umbrales de ruido: %v", err)
	}
	if float64(data.Noise) < noiseUmbralBajo {
		_, err = tx.Exec("INSERT INTO alerta (tipo, descripcion, valor_detectado, umbral, sala_id) VALUES ($1, $2, $3, $4, $5)",
			"critico", "Ruido por debajo del umbral", data.Noise, "bajo", salaID)
		if err != nil {
			return fmt.Errorf("error al crear alerta de ruido bajo: %v", err)
		}
	} else if float64(data.Noise) > noiseUmbralAlto {
		_, err = tx.Exec("INSERT INTO alerta (tipo, descripcion, valor_detectado, umbral, sala_id) VALUES ($1, $2, $3, $4, $5)",
			"critico", "Nivel de ruido crítico", data.Noise, "alto", salaID)
		if err != nil {
			return fmt.Errorf("error al crear alerta de ruido: %v", err)
		}
	} else if float64(data.Noise) >= noiseUmbralBajoPrev && float64(data.Noise) <= noiseUmbralAltoPrev {
		_, err = tx.Exec("INSERT INTO alerta (tipo, descripcion, valor_detectado, umbral, sala_id) VALUES ($1, $2, $3, $4, $5)",
			"preventivo", "Ruido por encima del umbral", data.Noise, "alto", salaID)
		if err != nil {
			return fmt.Errorf("error al crear alerta de ruido alta: %v", err)
		}
	}

	// Iluminación
	var lightUmbralBajo, lightUmbralAlto, lightUmbralBajoPrev, lightUmbralAltoPrev float64
	err = tx.QueryRow("SELECT umbral_bajo, umbral_alto, umbral_bajo_prev, umbral_alto_prev FROM umbrales WHERE sala_id = $1 AND variable_id = $2", salaID, lightVarID).Scan(&lightUmbralBajo, &lightUmbralAlto, &lightUmbralBajoPrev, &lightUmbralAltoPrev)
	if err != nil {
		return fmt.Errorf("error al obtener umbrales de iluminación: %v", err)
	}
	if float64(data.Light) < lightUmbralBajo {
		_, err = tx.Exec("INSERT INTO alerta (tipo, descripcion, valor_detectado, umbral, sala_id) VALUES ($1, $2, $3, $4, $5)",
			"critico", "Iluminación por debajo del umbral", data.Light, "bajo", salaID)
		if err != nil {
			return fmt.Errorf("error al crear alerta de iluminación baja: %v", err)
		}
	} else if float64(data.Light) > lightUmbralAlto {
		_, err = tx.Exec("INSERT INTO alerta (tipo, descripcion, valor_detectado, umbral, sala_id) VALUES ($1, $2, $3, $4, $5)",
			"critico", "Iluminación por encima del umbral", data.Light, "alto", salaID)
		if err != nil {
			return fmt.Errorf("error al crear alerta de iluminación alta: %v", err)
		}
	} else if (float64(data.Light) <= lightUmbralBajoPrev && float64(data.Light) >= lightUmbralBajo) || (float64(data.Light) >= lightUmbralAltoPrev && float64(data.Light) <= lightUmbralAlto) {
		_, err = tx.Exec("INSERT INTO alerta (tipo, descripcion, valor_detectado, umbral, sala_id) VALUES ($1, $2, $3, $4, $5)",
			"preventivo", "Iluminación transgrede el umbral", data.Light, "alto", salaID)
		if err != nil {
			return fmt.Errorf("error al crear alerta de iluminación alta: %v", err)
		}
	}

	return nil
}
