package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/models"

	"github.com/gorilla/mux"
)

type MedicionController struct {
	DB *sql.DB
}

func NewMedicionController(db *sql.DB) *MedicionController {
	return &MedicionController{DB: db}
}

func (c *MedicionController) Create(w http.ResponseWriter, r *http.Request) {
	var medicion models.Medicion
	if err := json.NewDecoder(r.Body).Decode(&medicion); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO medicion (fecha, valor, sensor_id) 
			  VALUES ($1, $2, $3) RETURNING id`
	err := c.DB.QueryRow(query, medicion.Fecha, medicion.Valor,
		medicion.SensorID).Scan(&medicion.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(medicion)
}

func (c *MedicionController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var medicion models.Medicion
	query := `SELECT id, fecha, valor, sensor_id 
			  FROM medicion WHERE id = $1`
	err = c.DB.QueryRow(query, id).Scan(&medicion.ID, &medicion.Fecha,
		&medicion.Valor, &medicion.SensorID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Medición no encontrada", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(medicion)
}

func (c *MedicionController) GetAll(w http.ResponseWriter, r *http.Request) {
	rows, err := c.DB.Query("SELECT id, fecha, valor, sensor_id FROM medicion")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var mediciones []models.Medicion
	for rows.Next() {
		var medicion models.Medicion
		if err := rows.Scan(&medicion.ID, &medicion.Fecha,
			&medicion.Valor, &medicion.SensorID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		mediciones = append(mediciones, medicion)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mediciones)
}

func (c *MedicionController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var medicion models.Medicion
	if err := json.NewDecoder(r.Body).Decode(&medicion); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `UPDATE medicion 
			  SET fecha = $1, valor = $2, sensor_id = $3 
			  WHERE id = $4`
	result, err := c.DB.Exec(query, medicion.Fecha, medicion.Valor,
		medicion.SensorID, id)
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
		http.Error(w, "Medición no encontrada", http.StatusNotFound)
		return
	}

	medicion.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(medicion)
}

func (c *MedicionController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	result, err := c.DB.Exec("DELETE FROM medicion WHERE id = $1", id)
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
		http.Error(w, "Medición no encontrada", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
