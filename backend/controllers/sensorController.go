package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/models"

	"github.com/gorilla/mux"
)

type SensorController struct {
	DB *sql.DB
}

func NewSensorController(db *sql.DB) *SensorController {
	return &SensorController{DB: db}
}

func (c *SensorController) Create(w http.ResponseWriter, r *http.Request) {
	var sensor models.Sensor
	if err := json.NewDecoder(r.Body).Decode(&sensor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO sensor (modelo, estado, fecha_instalacion, sala_id, variable_id) 
			  VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := c.DB.QueryRow(query, sensor.Modelo, sensor.Estado,
		sensor.FechaInstalacion, sensor.SalaID, sensor.VariableID).Scan(&sensor.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sensor)
}

func (c *SensorController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var sensor models.Sensor
	query := `SELECT id, modelo, estado, fecha_instalacion, sala_id, variable_id 
			  FROM sensor WHERE id = $1`
	err = c.DB.QueryRow(query, id).Scan(&sensor.ID, &sensor.Modelo, &sensor.Estado,
		&sensor.FechaInstalacion, &sensor.SalaID, &sensor.VariableID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Sensor no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sensor)
}

func (c *SensorController) GetAll(w http.ResponseWriter, r *http.Request) {
	rows, err := c.DB.Query("SELECT id, modelo, estado, fecha_instalacion, sala_id, variable_id FROM sensor")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var sensores []models.Sensor
	for rows.Next() {
		var sensor models.Sensor
		if err := rows.Scan(&sensor.ID, &sensor.Modelo, &sensor.Estado,
			&sensor.FechaInstalacion, &sensor.SalaID, &sensor.VariableID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sensores = append(sensores, sensor)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sensores)
}

func (c *SensorController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var sensor models.Sensor
	if err := json.NewDecoder(r.Body).Decode(&sensor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `UPDATE sensor 
			  SET modelo = $1, estado = $2, fecha_instalacion = $3, sala_id = $4, variable_id = $5 
			  WHERE id = $6`
	result, err := c.DB.Exec(query, sensor.Modelo, sensor.Estado,
		sensor.FechaInstalacion, sensor.SalaID, sensor.VariableID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Sensor no encontrado", http.StatusNotFound)
		return
	}

	sensor.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sensor)
}

func (c *SensorController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	result, err := c.DB.Exec("DELETE FROM sensor WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Sensor no encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
