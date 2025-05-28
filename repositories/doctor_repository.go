package repositories

import (
	"github.com/yourusername/fe/config"
	"github.com/yourusername/fe/models"
)

// DoctorRepository 医生信息表数据访问接口
type DoctorRepository interface {
	Create(doctor *models.Doctor) error
	FindByID(id uint) (*models.Doctor, error)
	FindAll(params map[string]interface{}, page, pageSize int) ([]models.Doctor, int64, error)
	Update(doctor *models.Doctor) error
	Delete(id uint) error
}

// doctorRepository 医生信息表数据访问实现
type doctorRepository struct{}

// NewDoctorRepository 创建医生信息表仓库实例
func NewDoctorRepository() DoctorRepository {
	return &doctorRepository{}
}

// Create 创建医生记录
func (r *doctorRepository) Create(doctor *models.Doctor) error {
	return config.DB.Create(doctor).Error
}

// FindByID 通过ID查找医生记录
func (r *doctorRepository) FindByID(id uint) (*models.Doctor, error) {
	var doctor models.Doctor
	err := config.DB.Preload("Department").First(&doctor, id).Error
	return &doctor, err
}

// FindAll 查找所有医生记录（支持筛选和分页）
func (r *doctorRepository) FindAll(params map[string]interface{}, page, pageSize int) ([]models.Doctor, int64, error) {
	var doctors []models.Doctor
	var total int64

	query := config.DB.Model(&models.Doctor{}).Preload("Department")

	// 应用筛选条件
	if params != nil {
		if name, ok := params["name"].(string); ok && name != "" {
			query = query.Where("name LIKE ?", "%"+name+"%")
		}
		if department, ok := params["department"].(string); ok && department != "" {
			query = query.Where("department LIKE ?", "%"+department+"%")
		}
	}

	// 获取总记录数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Find(&doctors).Error

	return doctors, total, err
}

// Update 更新医生记录
func (r *doctorRepository) Update(doctor *models.Doctor) error {
	return config.DB.Save(doctor).Error
}

// Delete 删除医生记录
func (r *doctorRepository) Delete(id uint) error {
	return config.DB.Delete(&models.Doctor{}, id).Error
}
