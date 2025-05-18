package services

import (
	"github.com/yourusername/fe/models"
	"github.com/yourusername/fe/repositories"
)

// MedicalExaminationService 病历检查表服务接口
type MedicalExaminationService interface {
	CreateExamination(exam *models.MedicalExamination) error
	GetExaminationByID(id uint) (*models.MedicalExamination, error)
	GetAllExaminations(params map[string]interface{}, page, pageSize int) ([]models.MedicalExamination, int64, error)
	UpdateExamination(exam *models.MedicalExamination) error
	DeleteExamination(id uint) error
	ExportExaminations(params map[string]interface{}) ([]models.MedicalExamination, error)
}

// medicalExaminationService 病历检查表服务实现
type medicalExaminationService struct {
	examRepo repositories.MedicalExaminationRepository
}

// NewMedicalExaminationService 创建病历检查表服务实例
func NewMedicalExaminationService(examRepo repositories.MedicalExaminationRepository) MedicalExaminationService {
	return &medicalExaminationService{
		examRepo: examRepo,
	}
}

// CreateExamination 创建病历检查记录
func (s *medicalExaminationService) CreateExamination(exam *models.MedicalExamination) error {
	return s.examRepo.Create(exam)
}

// GetExaminationByID 获取病历检查记录
func (s *medicalExaminationService) GetExaminationByID(id uint) (*models.MedicalExamination, error) {
	return s.examRepo.FindByID(id)
}

// GetAllExaminations 获取所有病历检查记录
func (s *medicalExaminationService) GetAllExaminations(params map[string]interface{}, page, pageSize int) ([]models.MedicalExamination, int64, error) {
	return s.examRepo.FindAll(params, page, pageSize)
}

// UpdateExamination 更新病历检查记录
func (s *medicalExaminationService) UpdateExamination(exam *models.MedicalExamination) error {
	return s.examRepo.Update(exam)
}

// DeleteExamination 删除病历检查记录
func (s *medicalExaminationService) DeleteExamination(id uint) error {
	return s.examRepo.Delete(id)
}

// ExportExaminations 导出病历检查记录
func (s *medicalExaminationService) ExportExaminations(params map[string]interface{}) ([]models.MedicalExamination, error) {
	return s.examRepo.ExportAll(params)
}
