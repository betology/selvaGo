package api

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Nombre struct {
	NombreID    int     `json:"NombreID"`
	FamiliaID   int     `json:"FamiliaID"`
	Nombre      string  `json:"Nombre"`
	Fecha       string  `json:"Fecha"`
	ProveedorID int     `json:"ProveedorID"`
	Precio      float64 `json:"Precio"`
	Inactivo    bool    `json:"Inactivo"`
}

type APIHandler struct {
	DB *sql.DB
}

func NewAPIHandler(db *sql.DB) *APIHandler {
	return &APIHandler{DB: db}
}

func (h *APIHandler) CreateNombre(c *gin.Context) {
	var nombre Nombre
	if err := c.ShouldBindJSON(&nombre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.DB.Exec("INSERT INTO Nombres (FamiliaID, Nombre, Fecha, ProveedorID, Precio, Inactivo) VALUES (?, ?, ?, ?, ?, ?)",
		nombre.FamiliaID, nombre.Nombre, nombre.Fecha, nombre.ProveedorID, nombre.Precio, nombre.Inactivo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	nombre.NombreID = int(id)
	c.JSON(http.StatusCreated, nombre)
}

func (h *APIHandler) GetNombres(c *gin.Context) {
	rows, err := h.DB.Query("SELECT NombreID, FamiliaID, Nombre, Fecha, ProveedorID, Precio, Inactivo FROM Nombres")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var nombres []Nombre
	for rows.Next() {
		var nombre Nombre
		if err := rows.Scan(&nombre.NombreID, &nombre.FamiliaID, &nombre.Nombre, &nombre.Fecha, &nombre.ProveedorID, &nombre.Precio, &nombre.Inactivo); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		nombres = append(nombres, nombre)
	}

	c.JSON(http.StatusOK, nombres)
}

func (h *APIHandler) GetNombreByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var nombre Nombre
	err = h.DB.QueryRow("SELECT NombreID, FamiliaID, Nombre, Fecha, ProveedorID, Precio, Inactivo FROM Nombres WHERE NombreID = ?", id).Scan(
		&nombre.NombreID, &nombre.FamiliaID, &nombre.Nombre, &nombre.Fecha, &nombre.ProveedorID, &nombre.Precio, &nombre.Inactivo)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Nombre not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, nombre)
}

func (h *APIHandler) UpdateNombre(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var nombre Nombre
	if err := c.ShouldBindJSON(&nombre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.DB.Exec("UPDATE Nombres SET FamiliaID = ?, Nombre = ?, Fecha = ?, ProveedorID = ?, Precio = ?, Inactivo = ? WHERE NombreID = ?",
		nombre.FamiliaID, nombre.Nombre, nombre.Fecha, nombre.ProveedorID, nombre.Precio, nombre.Inactivo, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	nombre.NombreID = id
	c.JSON(http.StatusOK, nombre)
}

func (h *APIHandler) SearchNombres(c *gin.Context) {
	nombre := c.Query("nombre")
	if nombre == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nombre parameter is required"})
		return
	}

	rows, err := h.DB.Query("SELECT NombreID, FamiliaID, Nombre, Fecha, ProveedorID, Precio, Inactivo FROM Nombres WHERE Nombre = ?", nombre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var nombres []Nombre
	for rows.Next() {
		var nombre Nombre
		if err := rows.Scan(&nombre.NombreID, &nombre.FamiliaID, &nombre.Nombre, &nombre.Fecha, &nombre.ProveedorID, &nombre.Precio, &nombre.Inactivo); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		nombres = append(nombres, nombre)
	}

	c.JSON(http.StatusOK, nombres)
}

func (h *APIHandler) DeleteNombre(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	_, err = h.DB.Exec("DELETE FROM Nombres WHERE NombreID = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Nombre deleted"})
}
