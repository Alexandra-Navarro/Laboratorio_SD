package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/models"
	"github.com/gorilla/mux"
)

type VariableAmbientalController struct {
	DB *sql.DB
}

func NewVariableAmbientalController(db *sql.DB) *VariableAmbientalController {
	return &VariableAmbientalController{DB: db}
}

func (c *VariableAmbientalController) Create(w http.ResponseWriter, r *http.Request) {
	var variable models.VariableAmbiental
	if err := json.NewDecoder(r.Body).Decode(&variable); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO variable_ambiental (nombre, unidad_medida, umbral_bajo, umbral_alto) 
			  VALUES ($1, $2, $3, $4) RETURNING id`
	err := c.DB.QueryRow(query, variable.Nombre, variable.UnidadMedida,
		variable.UmbralBajo, variable.UmbralAlto).Scan(&variable.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(variable)
}

func (c *VariableAmbientalController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var variable models.VariableAmbiental
	query := `SELECT id, nombre, unidad_medida, umbral_bajo, umbral_alto 
			  FROM variable_ambiental WHERE id = $1`
	err = c.DB.QueryRow(query, id).Scan(&variable.ID, &variable.Nombre,
		&variable.UnidadMedida, &variable.UmbralBajo, &variable.UmbralAlto)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Variable ambiental no encontrada", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(variable)
}

func (c *VariableAmbientalController) GetAll(w http.ResponseWriter, r *http.Request) {
	rows, err := c.DB.Query("SELECT id, nombre, unidad_medida, umbral_bajo, umbral_alto FROM variable_ambiental")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var variables []models.VariableAmbiental
	for rows.Next() {
		var variable models.VariableAmbiental
		if err := rows.Scan(&variable.ID, &variable.Nombre, &variable.UnidadMedida,
			&variable.UmbralBajo, &variable.UmbralAlto); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		variables = append(variables, variable)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(variables)
}

func (c *VariableAmbientalController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var variable models.VariableAmbiental
	if err := json.NewDecoder(r.Body).Decode(&variable); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `UPDATE variable_ambiental 
			  SET nombre = $1, unidad_medida = $2, umbral_bajo = $3, umbral_alto = $4 
			  WHERE id = $5`
	result, err := c.DB.Exec(query, variable.Nombre, variable.UnidadMedida,
		variable.UmbralBajo, variable.UmbralAlto, id)
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
		http.Error(w, "Variable ambiental no encontrada", http.StatusNotFound)
		return
	}

	variable.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(variable)
}

func (c *VariableAmbientalController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	result, err := c.DB.Exec("DELETE FROM variable_ambiental WHERE id = $1", id)
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
		http.Error(w, "Variable ambiental no encontrada", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
