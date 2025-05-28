package repositories

import (
	"github.com/yourusername/fe/config"
	"github.com/yourusername/fe/models"
)

type DepartmentRepository interface {
	Create(dept *models.Department) error
	FindByID(id uint) (*models.Department, error)
	FindAll() ([]models.Department, error)
	Update(dept *models.Department) error
	Delete(id uint) error
}

type departmentRepository struct{}

func NewDepartmentRepository() DepartmentRepository {
	return &departmentRepository{}
}

func (r *departmentRepository) Create(dept *models.Department) error {
	return config.DB.Create(dept).Error
}

func (r *departmentRepository) FindByID(id uint) (*models.Department, error) {
	var dept models.Department
	err := config.DB.First(&dept, id).Error
	return &dept, err
}

func (r *departmentRepository) FindAll() ([]models.Department, error) {
	var depts []models.Department
	err := config.DB.Find(&depts).Error
	return depts, err
}

func (r *departmentRepository) Update(dept *models.Department) error {
	return config.DB.Save(dept).Error
}

func (r *departmentRepository) Delete(id uint) error {
	return config.DB.Delete(&models.Department{}, id).Error
}
