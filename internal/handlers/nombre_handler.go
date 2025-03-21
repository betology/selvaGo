package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/betology/selvaGo/internal/models"
	"github.com/gin-gonic/gin"
)

// NombreHandler handles HTTP requests related to Nombres.
type NombreHandler struct {
	DB *sql.DB
}

// NewNombreHandler creates a new NombreHandler.
func NewNombreHandler(db *sql.DB) *NombreHandler {
	return &NombreHandler{DB: db}
}

// GetNombres retrieves all nombres.
func (h *NombreHandler) GetNombres(c *gin.Context) {
	rows, err := h.DB.Query("SELECT NombreID, FamiliaID, Nombre, Fecha, ProveedorID, Precio, Inactivo FROM Nombres")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var nombres []models.Nombre
	for rows.Next() {
		var n models.Nombre
		var fechaStr string
		if err := rows.Scan(&n.NombreID, &n.FamiliaID, &n.Nombre, &fechaStr, &n.ProveedorID, &n.Precio, &n.Inactivo); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		n.Fecha, err = time.Parse("2006-01-02", fechaStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		nombres = append(nombres, n)
	}

	c.JSON(http.StatusOK, nombres)
}

// GetNombreByID retrieves a nombre by its ID.
func (h *NombreHandler) GetNombreByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var n models.Nombre
	var fechaStr string
	err = h.DB.QueryRow("SELECT NombreID, FamiliaID, Nombre, Fecha, ProveedorID, Precio, Inactivo FROM Nombres WHERE NombreID = ?", id).Scan(&n.NombreID, &n.FamiliaID, &n.Nombre, &fechaStr, &n.ProveedorID, &n.Precio, &n.Inactivo)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Nombre not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	n.Fecha, err = time.Parse("2006-01-02", fechaStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, n)
}

// CreateNombre creates a new nombre.
func (h *NombreHandler) CreateNombre(c *gin.Context) {
	var n models.Nombre
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.DB.Exec("INSERT INTO Nombres (FamiliaID, Nombre, Fecha, ProveedorID, Precio, Inactivo) VALUES (?, ?, ?, ?, ?, ?)",
		n.FamiliaID, n.Nombre, n.Fecha.Format("2006-01-02"), n.ProveedorID, n.Precio, n.Inactivo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	n.NombreID = int(id)
	c.JSON(http.StatusCreated, n)
}

// UpdateNombre updates an existing nombre.
func (h *NombreHandler) UpdateNombre(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var n models.Nombre
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.DB.Exec("UPDATE Nombres SET FamiliaID = ?, Nombre = ?, Fecha = ?, ProveedorID = ?, Precio = ?, Inactivo = ? WHERE NombreID = ?",
		n.FamiliaID, n.Nombre, n.Fecha.Format("2006-01-02"), n.ProveedorID, n.Precio, n.Inactivo, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	n.NombreID = id
	c.JSON(http.StatusOK, n)
}

// DeleteNombre deletes a nombre.
func (h *NombreHandler) DeleteNombre(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
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
