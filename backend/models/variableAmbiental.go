package models

type VariableAmbiental struct {
	ID           int    `gorm:"primaryKey; autoIncrement" json:"id"`
	Nombre       string `json:"nombre"`
	UnidadMedida string `json:"unidad_medida"`
}

func (VariableAmbiental) TableName() string {
	return "variable_ambiental"
}
