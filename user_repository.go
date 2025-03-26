// repository/user_repository.go
package repository

import (
    "github.com/shinibufahaha/fengge/models"
    "github.com/jinzhu/gorm"
)

// UserRepository describes the database operations for users.
type UserRepository interface {
    Create(user *models.User) error
    GetByID(id uint) (*models.User, error)
    Update(user *models.User) error
    Delete(user *models.User) error
    GetAll() ([]models.User, error)
}

type userRepository struct {
    db *gorm.DB
}

// NewUserRepository instantiates a new UserRepository.
func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db}
}

func (r *userRepository) Create(user *models.User) error {
    return r.db.Create(user).Error
}

func (r *userRepository) GetByID(id uint) (*models.User, error) {
    var user models.User
    if err := r.db.First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) Update(user *models.User) error {
    return r.db.Save(user).Error
}

func (r *userRepository) Delete(user *models.User) error {
    return r.db.Delete(user).Error
}

func (r *userRepository) GetAll() ([]models.User, error) {
    var users []models.User
    if err := r.db.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}
