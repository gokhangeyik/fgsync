package render

import (
	"fmt"
	"gsync/internal/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ClientMainPage handles rendering the main page for the Tsunami UDP client.
// It uses the client_index.tmpl template with a title context parameter.
func ClientMainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "pages/client_index.tmpl", gin.H{
		"title": "Gsync - Tsunami UDP Client",
	})
}

// ServerMainPage handles rendering the main page for the Tsunami UDP server.
// It uses the server_index.tmpl template with a title context parameter.
func ServerMainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "pages/server_index.tmpl", gin.H{
		"title": "Gsync - Tsunami UDP Server",
	})
}

// FileListPartial handles rendering the file list partial template.
// It retrieves the list of files from a specified directory and passes them to the template.
// If there's an error accessing the directory, it renders an error message instead.
func FileListPartial(c *gin.Context) {
	path := "/home/gokko/Projects/gsync-go"
	files, err := tools.ListDir(path)
	if err != nil {
		c.HTML(http.StatusOK, "partials/filelist.tmpl", gin.H{
			"error": fmt.Sprintf("%s is not accessible!", path),
		})
		return
	}
	c.HTML(http.StatusOK, "partials/filelist.tmpl", gin.H{
		"fileList": files,
	})
}
