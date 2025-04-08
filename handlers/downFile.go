package handlers

import (
    "github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

func DownloadAPK(c *gin.Context) {
	filePath := "./uploads/sample.apk"

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(404, gin.H{"error": "File not found"})
		return
	}

    c.Header("Content-Type", "application/vnd.android.package-archive")
    c.Header("Content-Disposition", "attachment; filename="+filepath.Base(filePath))
    
    // Serve the file
    c.File(filePath)
}