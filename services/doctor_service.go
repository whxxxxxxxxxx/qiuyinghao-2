package services

import (
	"github.com/yourusername/fe/models"
	"github.com/yourusername/fe/repositories"
)

// DoctorService 医生信息表服务接口
type DoctorService interface {
	CreateDoctor(doctor *models.Doctor) error
	GetDoctorByID(id uint) (*models.Doctor, error)
	GetAllDoctors(params map[string]interface{}, page, pageSize int) ([]models.Doctor, int64, error)
	UpdateDoctor(doctor *models.Doctor) error
	DeleteDoctor(id uint) error
}

// doctorService 医生信息表服务实现
type doctorService struct {
	doctorRepo repositories.DoctorRepository
}

// NewDoctorService 创建医生信息表服务实例
func NewDoctorService(doctorRepo repositories.DoctorRepository) DoctorService {
	return &doctorService{
		doctorRepo: doctorRepo,
	}
}

// CreateDoctor 创建医生记录
func (s *doctorService) CreateDoctor(doctor *models.Doctor) error {
	return s.doctorRepo.Create(doctor)
}

// GetDoctorByID 获取医生记录
func (s *doctorService) GetDoctorByID(id uint) (*models.Doctor, error) {
	return s.doctorRepo.FindByID(id)
}

// GetAllDoctors 获取所有医生记录
func (s *doctorService) GetAllDoctors(params map[string]interface{}, page, pageSize int) ([]models.Doctor, int64, error) {
	return s.doctorRepo.FindAll(params, page, pageSize)
}

// UpdateDoctor 更新医生记录
func (s *doctorService) UpdateDoctor(doctor *models.Doctor) error {
	return s.doctorRepo.Update(doctor)
}

// DeleteDoctor 删除医生记录
func (s *doctorService) DeleteDoctor(id uint) error {
	return s.doctorRepo.Delete(id)
}
