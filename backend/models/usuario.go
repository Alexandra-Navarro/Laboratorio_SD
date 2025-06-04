package models

type Usuario struct {
	RutUsuario string `gorm:"primaryKey" json:"rut_usuario"`
	Nombre     string `json:"nombre"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Rol        string `json:"rol"`
	EscuelaID  int    `json:"escuela_id"`
}

func (Usuario) TableName() string {
	return "usuario"
}
