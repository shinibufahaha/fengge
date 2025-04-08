package handlers

import (
    "github.com/gin-gonic/gin"
	"os"
	"time"
	"log"
	"path/filepath"
	"gorm.io/gorm" 
	"fengge/utils"
	"fengge/db"
	"fengge/models"
)

func UploadPageHandler(c *gin.Context) {
    c.HTML(200, "upload.html", gin.H{
        "title": "Upload APK File",
    })
}

func UploadFileHandler(c *gin.Context) {
    cfg := utils.LoadConfig("config.yaml")
    file, err := c.FormFile("file")
    if err != nil {
        log.Printf("Error getting file: %v", err)
        c.HTML(400, "upload.html", gin.H{
            "title": "Upload APK File",
            "error": "Failed to get file: " + err.Error(),
        })
        return
    }
    log.Printf("Received file upload request for: %s", file.Filename)

    timestamp := time.Now().Format("20060102150405")
    savedFilePath := filepath.Join("uploads", timestamp + "_" + file.Filename)
    if err := c.SaveUploadedFile(file, savedFilePath); err != nil {
        log.Printf("Error saving file: %v", err)
        c.HTML(500, "upload.html", gin.H{
            "title": "Upload APK File",
            "error": "Failed to save file: " + err.Error(),
        })
        return
    }
    log.Printf("File saved successfully at: %s", savedFilePath)

    hash, err := utils.GenerateFileHash(savedFilePath)
    if err != nil {
        log.Printf("Error generating hash: %v", err)
        c.HTML(500, "upload.html", gin.H{
            "title": "Upload APK File",
            "error": "Failed to generate hash: " + err.Error(),
        })
        return
    }
    log.Printf("Generated hash for file: %s", hash)

    var existingFile models.File
    result := db.GetDB().First(&existingFile, "hash_value = ?", hash)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            log.Printf("New file detected, creating database entry")
            cpgName := file.Filename[:len(file.Filename)-len(filepath.Ext(file.Filename))] + ".cpg"
            cpgPath := filepath.Join(cfg.Path.ProjectPath, "cpgs", cpgName)
            newFile := models.File{HashValue: hash, FilePath: savedFilePath, CpgPath: cpgPath}
            if err := db.GetDB().Create(&newFile).Error; err != nil {
                log.Printf("Error creating database entry: %v", err)
                c.JSON(500, gin.H{"error": "Database error: " + err.Error()})
                return
            }

            log.Printf("Starting CPG generation for file: %s", savedFilePath)
            utils.GenerateCPG(savedFilePath)
            log.Printf("CPG generation completed")
            c.JSON(200, gin.H{"message": "File uploaded and project created!"})
        } else {
            log.Printf("Database error: %v", result.Error)
            c.JSON(500, gin.H{"error": "Database error: " + result.Error.Error()})
        }
    } else {
        log.Printf("File with hash %s already exists", hash)
        c.JSON(200, gin.H{"message": "File already exists. No new project created."})
        // Remove the uploaded file since it already exists in the database
        if err := os.Remove(savedFilePath); err != nil {
            log.Printf("Error removing duplicate file: %v", err)
            c.JSON(500, gin.H{"error": "Failed to remove duplicate file: " + err.Error()})
            return
        }
        log.Printf("Duplicate file removed successfully")
    }

    c.HTML(200, "upload.html", gin.H{
        "title": "Upload APK File",
        "success": "File uploaded successfully!",
    })
}