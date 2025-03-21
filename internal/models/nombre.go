package models

import "time"

// Nombre represents a row in the Nombres table.
type Nombre struct {
	NombreID    int       `json:"nombre_id"`
	FamiliaID   int       `json:"familia_id"`
	Nombre      string    `json:"nombre"`
	Fecha       time.Time `json:"fecha"`
	ProveedorID int       `json:"proveedor_id"`
	Precio      float64   `json:"precio"`
	Inactivo    bool      `json:"inactivo"`
}
