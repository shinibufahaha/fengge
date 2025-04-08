// main.go
package main

import (
    "github.com/gin-gonic/gin"
    "fengge/handlers"
    "fengge/db"
    //"fengge/utils"
)



func main() {
    // config := utils.LoadConfig("config.yaml")

    _ = db.GetDB()
    // Set up the Gin router
    router := gin.Default()

    // Serve static files and templates
    router.LoadHTMLGlob("templates/*")
    router.Static("/static", "./static")

    
    router.GET("/", handlers.DefaultHandler)
    // Add route for the upload page
    router.GET("/upload", handlers.UploadPageHandler)

    // Add route for handling file upload
    router.POST("/upload", handlers.UploadFileHandler)

    router.GET("/download", handlers.DownloadAPK)

    // Start the server on port 8080.
    router.Run(":8080")
}
