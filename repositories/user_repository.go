package repositories

import (
	"github.com/yourusername/fe/config"
	"github.com/yourusername/fe/models"
)

// UserRepository 用户数据访问接口
type UserRepository interface {
	Create(user *models.User) error
	FindByID(id uint) (*models.User, error)
	FindAll() ([]models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
}

// userRepository 用户数据访问实现
type userRepository struct{}

// NewUserRepository 创建用户仓库实例
func NewUserRepository() UserRepository {
	return &userRepository{}
}

// Create 创建用户
func (r *userRepository) Create(user *models.User) error {
	return config.DB.Create(user).Error
}

// FindByID 通过ID查找用户
func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := config.DB.First(&user, id).Error
	return &user, err
}

// FindAll 查找所有用户
func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := config.DB.Find(&users).Error
	return users, err
}

// Update 更新用户
func (r *userRepository) Update(user *models.User) error {
	return config.DB.Save(user).Error
}

// Delete 删除用户
func (r *userRepository) Delete(id uint) error {
	return config.DB.Delete(&models.User{}, id).Error
}
