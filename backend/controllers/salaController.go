package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/models"
	"github.com/gorilla/mux"
)

type SalaController struct {
	DB *sql.DB
}

func NewSalaController(db *sql.DB) *SalaController {
	return &SalaController{DB: db}
}

func (c *SalaController) Create(w http.ResponseWriter, r *http.Request) {
	var sala models.Sala
	if err := json.NewDecoder(r.Body).Decode(&sala); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO sala (nombre, escuela_id) VALUES ($1, $2) RETURNING id`
	err := c.DB.QueryRow(query, sala.Nombre, sala.EscuelaID).Scan(&sala.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sala)
}

func (c *SalaController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var sala models.Sala
	query := `SELECT id, nombre, escuela_id FROM sala WHERE id = $1`
	err = c.DB.QueryRow(query, id).Scan(&sala.ID, &sala.Nombre, &sala.EscuelaID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Sala no encontrada", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sala)
}

func (c *SalaController) GetAll(w http.ResponseWriter, r *http.Request) {
	rows, err := c.DB.Query("SELECT id, nombre, escuela_id FROM sala")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var salas []models.Sala
	for rows.Next() {
		var sala models.Sala
		if err := rows.Scan(&sala.ID, &sala.Nombre, &sala.EscuelaID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		salas = append(salas, sala)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(salas)
}

func (c *SalaController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var sala models.Sala
	if err := json.NewDecoder(r.Body).Decode(&sala); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `UPDATE sala SET nombre = $1, escuela_id = $2 WHERE id = $3`
	result, err := c.DB.Exec(query, sala.Nombre, sala.EscuelaID, id)
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
		http.Error(w, "Sala no encontrada", http.StatusNotFound)
		return
	}

	sala.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sala)
}

func (c *SalaController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	result, err := c.DB.Exec("DELETE FROM sala WHERE id = $1", id)
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
		http.Error(w, "Sala no encontrada", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
