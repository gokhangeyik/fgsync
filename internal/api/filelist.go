// Package api provides HTTP API handlers for the gsync application.
package api

import (
	"net/http"

	"gsync/internal/tools"

	"github.com/gin-gonic/gin"
)

// ErrResponse represents a standard error response structure for API errors.
type ErrResponse struct {
	Message string `json:"message"`
}

// GetFileList handles HTTP requests to list files in the data directory.
// It returns a JSON list of files or an error if the directory is not accessible.
func GetFileList(c *gin.Context) {
	files, err := tools.ListDir("/home/gokko/Projects/gsync-go")
	if err != nil {
		c.JSON(http.StatusNotFound, ErrResponse{Message: "directory not accessible"})
		return
	}
	c.JSON(http.StatusOK, files)
}
