package handlers

// import (
// 	"fmt"
//     "github.com/gin-gonic/gin"
// 	"fengge/models"
// 	"fengge/db"
// 	"fengge/utils"
// )

// func Start(c *gin.Context, id int) {
// 	config := utils.LoadConfig("config.yaml")
// 	var existingProject models.Project
//     result := db.GetDB().First(&existingProject, "_id = ?", id)
// 	cmd := fmt.Sprintf("%s %s --script %s --param exports=%s", config.JoernPath.JoernParse, cpgPath, scriptPath, exportsPath)
// }