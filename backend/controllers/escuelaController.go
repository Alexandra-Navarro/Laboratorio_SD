package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/models"
	"github.com/gorilla/mux"
)

type EscuelaController struct {
	DB *sql.DB
}

func NewEscuelaController(db *sql.DB) *EscuelaController {
	return &EscuelaController{DB: db}
}

func (c *EscuelaController) Create(w http.ResponseWriter, r *http.Request) {
	var escuela models.Escuela
	if err := json.NewDecoder(r.Body).Decode(&escuela); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO escuela (nombre, direccion, comuna) VALUES ($1, $2, $3) RETURNING id`
	err := c.DB.QueryRow(query, escuela.Nombre, escuela.Direccion, escuela.Comuna).Scan(&escuela.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(escuela)
}

func (c *EscuelaController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var escuela models.Escuela
	query := `SELECT id, nombre, direccion, comuna FROM escuela WHERE id = $1`
	err = c.DB.QueryRow(query, id).Scan(&escuela.ID, &escuela.Nombre, &escuela.Direccion, &escuela.Comuna)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Escuela no encontrada", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(escuela)
}

func (c *EscuelaController) GetAll(w http.ResponseWriter, r *http.Request) {
	rows, err := c.DB.Query("SELECT id, nombre, direccion, comuna FROM escuela")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var escuelas []models.Escuela
	for rows.Next() {
		var escuela models.Escuela
		if err := rows.Scan(&escuela.ID, &escuela.Nombre, &escuela.Direccion, &escuela.Comuna); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		escuelas = append(escuelas, escuela)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(escuelas)
}

func (c *EscuelaController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var escuela models.Escuela
	if err := json.NewDecoder(r.Body).Decode(&escuela); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `UPDATE escuela SET nombre = $1, direccion = $2, comuna = $3 WHERE id = $4`
	result, err := c.DB.Exec(query, escuela.Nombre, escuela.Direccion, escuela.Comuna, id)
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
		http.Error(w, "Escuela no encontrada", http.StatusNotFound)
		return
	}

	escuela.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(escuela)
}

func (c *EscuelaController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	result, err := c.DB.Exec("DELETE FROM escuela WHERE id = $1", id)
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
		http.Error(w, "Escuela no encontrada", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
