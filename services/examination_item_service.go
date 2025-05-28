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
	CreateMaterial(material *models.Material) error
	GetMaterialByID(id uint) (*models.Material, error)
	GetAllMaterials() ([]models.Material, error)
	UpdateMaterial(material *models.Material) error
	DeleteMaterial(id uint) error
	AddMaterialToItem(itemID, materialID uint, quantity int) error
	RemoveMaterialFromItem(itemID, materialID uint) error
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

// CreateMaterial 创建材料记录
func (s *examinationItemService) CreateMaterial(material *models.Material) error {
	return s.itemRepo.CreateMaterial(material)
}

// GetMaterialByID 获取材料记录详情
func (s *examinationItemService) GetMaterialByID(id uint) (*models.Material, error) {
	return s.itemRepo.FindMaterialByID(id)
}

// GetAllMaterials 获取材料记录列表
func (s *examinationItemService) GetAllMaterials() ([]models.Material, error) {
	return s.itemRepo.FindAllMaterials()
}

// UpdateMaterial 更新材料记录
func (s *examinationItemService) UpdateMaterial(material *models.Material) error {
	return s.itemRepo.UpdateMaterial(material)
}

// DeleteMaterial 删除材料记录
func (s *examinationItemService) DeleteMaterial(id uint) error {
	return s.itemRepo.DeleteMaterial(id)
}

// AddMaterialToItem 向检查项目添加材料
func (s *examinationItemService) AddMaterialToItem(itemID, materialID uint, quantity int) error {
	return s.itemRepo.AddMaterialToItem(itemID, materialID, quantity)
}

// RemoveMaterialFromItem 从检查项目中移除材料
func (s *examinationItemService) RemoveMaterialFromItem(itemID, materialID uint) error {
	return s.itemRepo.RemoveMaterialFromItem(itemID, materialID)
}
