package models

type Escuela struct {
	ID        int    `gorm:"primaryKey; autoIncrement" json:"id"`
	Nombre    string `json:"nombre"`
	Direccion string `json:"direccion"`
	Comuna    string `json:"comuna"`
}

func (Escuela) TableName() string {
	return "escuela"
}
