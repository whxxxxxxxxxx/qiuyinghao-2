package repositories

import (
	"github.com/yourusername/fe/config"
	"github.com/yourusername/fe/models"
)

// ExaminationItemRepository 检查项目表数据访问接口
type ExaminationItemRepository interface {
	Create(item *models.ExaminationItem) error
	FindByID(id uint) (*models.ExaminationItem, error)
	FindAll(params map[string]interface{}, page, pageSize int) ([]models.ExaminationItem, int64, error)
	Update(item *models.ExaminationItem) error
	Delete(id uint) error
}

// examinationItemRepository 检查项目表数据访问实现
type examinationItemRepository struct{}

// NewExaminationItemRepository 创建检查项目表仓库实例
func NewExaminationItemRepository() ExaminationItemRepository {
	return &examinationItemRepository{}
}

// Create 创建检查项目记录
func (r *examinationItemRepository) Create(item *models.ExaminationItem) error {
	return config.DB.Create(item).Error
}

// FindByID 通过ID查找检查项目记录
func (r *examinationItemRepository) FindByID(id uint) (*models.ExaminationItem, error) {
	var item models.ExaminationItem
	err := config.DB.First(&item, id).Error
	return &item, err
}

// FindAll 查找所有检查项目记录（支持筛选和分页）
func (r *examinationItemRepository) FindAll(params map[string]interface{}, page, pageSize int) ([]models.ExaminationItem, int64, error) {
	var items []models.ExaminationItem
	var total int64

	query := config.DB.Model(&models.ExaminationItem{})

	// 应用筛选条件
	if params != nil {
		if name, ok := params["name"].(string); ok && name != "" {
			query = query.Where("name LIKE ?", "%"+name+"%")
		}
	}

	// 获取总记录数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Find(&items).Error

	return items, total, err
}

// Update 更新检查项目记录
func (r *examinationItemRepository) Update(item *models.ExaminationItem) error {
	return config.DB.Save(item).Error
}

// Delete 删除检查项目记录
func (r *examinationItemRepository) Delete(id uint) error {
	return config.DB.Delete(&models.ExaminationItem{}, id).Error
}
