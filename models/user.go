// models/user.go
package models

// User represents a user in the system.
type User struct {
    ID       uint   `json:"id" gorm:"primary_key"`
    Username string `json:"username" binding:"required"`
    Email    string `json:"email" binding:"required,email" gorm:"unique"`
    Password string `json:"password" binding:"required"`
}
