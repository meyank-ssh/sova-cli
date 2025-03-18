package routes

import (
	"github.com/gin-gonic/gin"
	"{{.ModuleName}}/internal/handlers"
	{{if .UseZap}}"{{.ModuleName}}/internal/middleware"{{end}}
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(router *gin.Engine) {
	{{if .UseZap}}// Add logging middleware
	router.Use(middleware.LoggingMiddleware())
	{{end}}

	// API routes
	api := router.Group("/api")
	{
		api.GET("/ping", handlers.PingHandler)
		api.GET("/health", handlers.HealthHandler)
	}
}