package services

import (
	"github.com/yourusername/fe/models"
	"github.com/yourusername/fe/repositories"
)

type DepartmentService interface {
	CreateDepartment(dept *models.Department) error
	GetDepartmentByID(id uint) (*models.Department, error)
	GetAllDepartments() ([]models.Department, error)
	UpdateDepartment(dept *models.Department) error
	DeleteDepartment(id uint) error
}

type departmentService struct {
	repo repositories.DepartmentRepository
}

func NewDepartmentService(repo repositories.DepartmentRepository) DepartmentService {
	return &departmentService{repo: repo}
}

func (s *departmentService) CreateDepartment(dept *models.Department) error {
	return s.repo.Create(dept)
}

func (s *departmentService) GetDepartmentByID(id uint) (*models.Department, error) {
	return s.repo.FindByID(id)
}

func (s *departmentService) GetAllDepartments() ([]models.Department, error) {
	return s.repo.FindAll()
}

func (s *departmentService) UpdateDepartment(dept *models.Department) error {
	return s.repo.Update(dept)
}

func (s *departmentService) DeleteDepartment(id uint) error {
	return s.repo.Delete(id)
}
