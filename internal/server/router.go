// Package server provides HTTP server functionality for the gsync application,
// including router configuration and server execution.
package server

import (
	api "gsync/internal/api"
	config "gsync/internal/config"
	render "gsync/web"

	"github.com/gin-gonic/gin"
)

var role = config.GetConfig().Role

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/**/*.tmpl")
	router.Static("/static", "web/static")

	// Frontend routes
	if role == "SERVER" {
		router.GET("/", render.ServerMainPage)
	} else {
		router.GET("/", render.ClientMainPage)
	}

	// Frontend partial routes
	router.GET("/partial/filelist", render.FileListPartial)

	// API routes
	apiRouter := router.Group("/api/v1")
	apiRouter.GET("/file/list", api.GetFileList)

	return router
}

func Run(addr string) error {
	gin.SetMode("release")
	router := SetupRouter()
	return router.Run(addr)
}
