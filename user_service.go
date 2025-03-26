// service/user_service.go
package service

import (
    "errors"
    "github.com/shinibufahaha/fengge/models"
    "github.com/shinibufahaha/fengge/repository"
	
)

type UserService interface {
    CreateUser(user *models.User) error
    GetUser(id uint) (*models.User, error)
    UpdateUser(user *models.User) error
    DeleteUser(id uint) error
    ListUsers() ([]models.User, error)
}

type userService struct {
    repo repository.UserRepository
}

// NewUserService returns a new UserService.
func NewUserService(repo repository.UserRepository) UserService {
    return &userService{repo}
}

func (s *userService) CreateUser(user *models.User) error {
    // Here, you can add business logic such as password hashing.
    return s.repo.Create(user)
}

func (s *userService) GetUser(id uint) (*models.User, error) {
    return s.repo.GetByID(id)
}

func (s *userService) UpdateUser(user *models.User) error {
    if user.ID == 0 {
        return errors.New("invalid user id")
    }
    // Additional business logic can go here.
    return s.repo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
    user, err := s.repo.GetByID(id)
    if err != nil {
        return err
    }
    return s.repo.Delete(user)
}

func (s *userService) ListUsers() ([]models.User, error) {
    return s.repo.GetAll()
}
