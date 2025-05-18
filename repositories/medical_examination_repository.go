package repositories

import (
	"github.com/yourusername/fe/config"
	"github.com/yourusername/fe/models"
)

// MedicalExaminationRepository 病历检查表数据访问接口
type MedicalExaminationRepository interface {
	Create(exam *models.MedicalExamination) error
	FindByID(id uint) (*models.MedicalExamination, error)
	FindAll(params map[string]interface{}, page, pageSize int) ([]models.MedicalExamination, int64, error)
	Update(exam *models.MedicalExamination) error
	Delete(id uint) error
	ExportAll(params map[string]interface{}) ([]models.MedicalExamination, error)
}

// medicalExaminationRepository 病历检查表数据访问实现
type medicalExaminationRepository struct{}

// NewMedicalExaminationRepository 创建病历检查表仓库实例
func NewMedicalExaminationRepository() MedicalExaminationRepository {
	return &medicalExaminationRepository{}
}

// Create 创建病历检查记录
func (r *medicalExaminationRepository) Create(exam *models.MedicalExamination) error {
	return config.DB.Create(exam).Error
}

// FindByID 通过ID查找病历检查记录
func (r *medicalExaminationRepository) FindByID(id uint) (*models.MedicalExamination, error) {
	var exam models.MedicalExamination
	err := config.DB.Preload("ExaminationItem").Preload("Doctor").First(&exam, id).Error
	return &exam, err
}

// FindAll 查找所有病历检查记录（支持筛选和分页）
func (r *medicalExaminationRepository) FindAll(params map[string]interface{}, page, pageSize int) ([]models.MedicalExamination, int64, error) {
	var exams []models.MedicalExamination
	var total int64

	query := config.DB.Model(&models.MedicalExamination{})

	// 应用筛选条件
	if params != nil {
		if patientName, ok := params["patientName"].(string); ok && patientName != "" {
			query = query.Where("patient_name LIKE ?", "%"+patientName+"%")
		}
		if doctorID, ok := params["doctor"].(uint); ok && doctorID > 0 {
			query = query.Where("doctor_id = ?", doctorID)
		}
		if examinationItemID, ok := params["examinationItem"].(uint); ok && examinationItemID > 0 {
			query = query.Where("examination_item_id = ?", examinationItemID)
		}
	}

	// 获取总记录数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Preload("ExaminationItem").Preload("Doctor").Find(&exams).Error

	return exams, total, err
}

// Update 更新病历检查记录
func (r *medicalExaminationRepository) Update(exam *models.MedicalExamination) error {
	return config.DB.Save(exam).Error
}

// Delete 删除病历检查记录
func (r *medicalExaminationRepository) Delete(id uint) error {
	return config.DB.Delete(&models.MedicalExamination{}, id).Error
}

// ExportAll 导出所有病历检查记录（支持筛选，不分页）
func (r *medicalExaminationRepository) ExportAll(params map[string]interface{}) ([]models.MedicalExamination, error) {
	var exams []models.MedicalExamination

	query := config.DB.Model(&models.MedicalExamination{}).Preload("ExaminationItem").Preload("Doctor")

	// 应用筛选条件
	if params != nil {
		if patientName, ok := params["patientName"].(string); ok && patientName != "" {
			query = query.Where("patient_name LIKE ?", "%"+patientName+"%")
		}
		if doctorID, ok := params["doctorID"].(uint); ok && doctorID > 0 {
			query = query.Where("doctor_id = ?", doctorID)
		}
		if examinationItemID, ok := params["examinationItemID"].(uint); ok && examinationItemID > 0 {
			query = query.Where("examination_item_id = ?", examinationItemID)
		}
	}

	// 查询所有符合条件的记录（不分页）
	err := query.Preload("ExaminationItem").Preload("Doctor").Find(&exams).Error

	return exams, err
}
