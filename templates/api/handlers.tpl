package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthHandler returns a 200 OK response if the service is healthy
func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"service":   "{{.ProjectName}}",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

// PingHandler returns a simple pong response
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// HomeHandler renders the home page
func HomeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":       "{{.ProjectName}}",
		"description": "{{.ProjectDescription}}",
	})
}

// ErrorHandler returns a JSON error response
func ErrorHandler(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"error": message,
	})
}

// NotFoundHandler handles 404 errors
func NotFoundHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": "Resource not found",
	})
}

// APIResponse is a standard API response structure
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SendSuccessResponse sends a success response
func SendSuccessResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// SendErrorResponse sends an error response
func SendErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, APIResponse{
		Success: false,
		Error:   message,
	})
} 