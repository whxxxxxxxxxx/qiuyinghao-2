package services

import (
	"github.com/yourusername/fe/models"
	"github.com/yourusername/fe/repositories"
)

// ExaminationItemService 检查项目表服务接口
type ExaminationItemService interface {
	CreateExaminationItem(item *models.ExaminationItem) error
	GetExaminationItemByID(id uint) (*models.ExaminationItem, error)
	GetAllExaminationItems(params map[string]interface{}, page, pageSize int) ([]models.ExaminationItem, int64, error)
	UpdateExaminationItem(item *models.ExaminationItem) error
	DeleteExaminationItem(id uint) error
}

// examinationItemService 检查项目表服务实现
type examinationItemService struct {
	itemRepo repositories.ExaminationItemRepository
}

// NewExaminationItemService 创建检查项目表服务实例
func NewExaminationItemService(itemRepo repositories.ExaminationItemRepository) ExaminationItemService {
	return &examinationItemService{
		itemRepo: itemRepo,
	}
}

// CreateExaminationItem 创建检查项目记录
func (s *examinationItemService) CreateExaminationItem(item *models.ExaminationItem) error {
	return s.itemRepo.Create(item)
}

// GetExaminationItemByID 获取检查项目记录
func (s *examinationItemService) GetExaminationItemByID(id uint) (*models.ExaminationItem, error) {
	return s.itemRepo.FindByID(id)
}

// GetAllExaminationItems 获取所有检查项目记录
func (s *examinationItemService) GetAllExaminationItems(params map[string]interface{}, page, pageSize int) ([]models.ExaminationItem, int64, error) {
	return s.itemRepo.FindAll(params, page, pageSize)
}

// UpdateExaminationItem 更新检查项目记录
func (s *examinationItemService) UpdateExaminationItem(item *models.ExaminationItem) error {
	return s.itemRepo.Update(item)
}

// DeleteExaminationItem 删除检查项目记录
func (s *examinationItemService) DeleteExaminationItem(id uint) error {
	return s.itemRepo.Delete(id)
}
