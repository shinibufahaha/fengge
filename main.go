// main.go
package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    "github.com/shinibufahaha/fengge/handler"
	"github.com/shinibufahaha/fengge/models"
    "github.com/shinibufahaha/fengge/repository"
    "github.com/shinibufahaha/fengge/service"
)

func main() {
    // Initialize database connection (using SQLite for simplicity)
    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Close()

    // Migrate the schema: automatically create the "users" table
    db.AutoMigrate(&models.User{})

    // Initialize repository, service, and handler layers
    userRepo := repository.NewUserRepository(db)
    userService := service.NewUserService(userRepo)
    userHandler := handler.NewUserHandler(userService)

    // Set up the Gin router
    router := gin.Default()

    // Define RESTful endpoints for user management.
    userRoutes := router.Group("/users")
    {
        userRoutes.POST("/", userHandler.CreateUser)
        userRoutes.GET("/", userHandler.GetUsers)
        userRoutes.GET("/:id", userHandler.GetUser)
        userRoutes.PUT("/:id", userHandler.UpdateUser)
        userRoutes.DELETE("/:id", userHandler.DeleteUser)
    }

    // Start the server on port 8080.
    router.Run(":8080")
}
