package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Alexandra-Navarro/Laboratorio_SD/backend/models"

	"github.com/gorilla/mux"
)

type UsuarioController struct {
	DB *sql.DB
}

func NewUsuarioController(db *sql.DB) *UsuarioController {
	return &UsuarioController{DB: db}
}

func (c *UsuarioController) Create(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO usuario (rut_usuario, nombre, email, password, rol, escuela_id) 
			  VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := c.DB.Exec(query, usuario.RutUsuario, usuario.Nombre,
		usuario.Email, usuario.Password, usuario.Rol, usuario.EscuelaID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usuario)
}

func (c *UsuarioController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rut := vars["rut"]

	var usuario models.Usuario
	query := `SELECT rut_usuario, nombre, email, password, rol, escuela_id 
			  FROM usuario WHERE rut_usuario = $1`
	err := c.DB.QueryRow(query, rut).Scan(&usuario.RutUsuario, &usuario.Nombre,
		&usuario.Email, &usuario.Password, &usuario.Rol, &usuario.EscuelaID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Usuario no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuario)
}

func (c *UsuarioController) GetAll(w http.ResponseWriter, r *http.Request) {
	rows, err := c.DB.Query("SELECT rut_usuario, nombre, email, password, rol, escuela_id FROM usuario")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var usuarios []models.Usuario
	for rows.Next() {
		var usuario models.Usuario
		if err := rows.Scan(&usuario.RutUsuario, &usuario.Nombre,
			&usuario.Email, &usuario.Password, &usuario.Rol, &usuario.EscuelaID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		usuarios = append(usuarios, usuario)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

func (c *UsuarioController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rut := vars["rut"]

	var usuario models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `UPDATE usuario 
			  SET nombre = $1, email = $2, password = $3, rol = $4, escuela_id = $5 
			  WHERE rut_usuario = $6`
	result, err := c.DB.Exec(query, usuario.Nombre, usuario.Email,
		usuario.Password, usuario.Rol, usuario.EscuelaID, rut)
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
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	usuario.RutUsuario = rut
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuario)
}

func (c *UsuarioController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rut := vars["rut"]

	result, err := c.DB.Exec("DELETE FROM usuario WHERE rut_usuario = $1", rut)
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
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
