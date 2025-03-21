package routes

import (
	"github.com/betology/selvaGo/internal/handlers"
	"github.com/gin-gonic/gin"
)

// SetupNombreRoutes sets up the routes for Nombres.
func SetupNombreRoutes(router *gin.Engine, handler *handlers.NombreHandler) {
	nombresGroup := router.Group("/nombres")
	{
		nombresGroup.GET("", handler.GetNombres)
		nombresGroup.GET("/:id", handler.GetNombreByID)
		nombresGroup.POST("", handler.CreateNombre)
		nombresGroup.PUT("/:id", handler.UpdateNombre)
		nombresGroup.DELETE("/:id", handler.DeleteNombre)
	}
}
