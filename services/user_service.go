package services

import (
	"github.com/yourusername/fe/models"
	"github.com/yourusername/fe/repositories"
)

// UserService 用户服务接口
type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}

// userService 用户服务实现
type userService struct {
	userRepo repositories.UserRepository
}

// NewUserService 创建用户服务实例
func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// CreateUser 创建用户
func (s *userService) CreateUser(user *models.User) error {
	return s.userRepo.Create(user)
}

// GetUserByID 获取用户
func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

// GetAllUsers 获取所有用户
func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.FindAll()
}

// UpdateUser 更新用户
func (s *userService) UpdateUser(user *models.User) error {
	return s.userRepo.Update(user)
}

// DeleteUser 删除用户
func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}
