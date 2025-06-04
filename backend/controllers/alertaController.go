package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/models"
	"github.com/gorilla/mux"
)

type AlertaController struct {
	DB *sql.DB
}

func NewAlertaController(db *sql.DB) *AlertaController {
	return &AlertaController{DB: db}
}

func (c *AlertaController) Create(w http.ResponseWriter, r *http.Request) {
	var alerta models.Alerta
	if err := json.NewDecoder(r.Body).Decode(&alerta); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO alerta (tipo, descripcion, valor_detectado, umbral, fecha, sala_id) 
			  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := c.DB.QueryRow(query, alerta.Tipo, alerta.Descripcion,
		alerta.ValorDetectado, alerta.Umbral, alerta.Fecha, alerta.SalaID).Scan(&alerta.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(alerta)
}

func (c *AlertaController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var alerta models.Alerta
	query := `SELECT id, tipo, descripcion, valor_detectado, umbral, fecha, sala_id 
			  FROM alerta WHERE id = $1`
	err = c.DB.QueryRow(query, id).Scan(&alerta.ID, &alerta.Tipo, &alerta.Descripcion,
		&alerta.ValorDetectado, &alerta.Umbral, &alerta.Fecha, &alerta.SalaID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Alerta no encontrada", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alerta)
}

func (c *AlertaController) GetAll(w http.ResponseWriter, r *http.Request) {
	rows, err := c.DB.Query("SELECT id, tipo, descripcion, valor_detectado, umbral, fecha, sala_id FROM alerta")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var alertas []models.Alerta
	for rows.Next() {
		var alerta models.Alerta
		if err := rows.Scan(&alerta.ID, &alerta.Tipo, &alerta.Descripcion,
			&alerta.ValorDetectado, &alerta.Umbral, &alerta.Fecha, &alerta.SalaID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		alertas = append(alertas, alerta)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alertas)
}

func (c *AlertaController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var alerta models.Alerta
	if err := json.NewDecoder(r.Body).Decode(&alerta); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `UPDATE alerta 
			  SET tipo = $1, descripcion = $2, valor_detectado = $3, umbral = $4, fecha = $5, sala_id = $6 
			  WHERE id = $7`
	result, err := c.DB.Exec(query, alerta.Tipo, alerta.Descripcion,
		alerta.ValorDetectado, alerta.Umbral, alerta.Fecha, alerta.SalaID, id)
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
		http.Error(w, "Alerta no encontrada", http.StatusNotFound)
		return
	}

	alerta.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alerta)
}

func (c *AlertaController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	result, err := c.DB.Exec("DELETE FROM alerta WHERE id = $1", id)
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
		http.Error(w, "Alerta no encontrada", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
