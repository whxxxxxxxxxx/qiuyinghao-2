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
	CreateMaterial(material *models.Material) error
	FindMaterialByID(id uint) (*models.Material, error)
	FindAllMaterials() ([]models.Material, error)
	UpdateMaterial(material *models.Material) error
	DeleteMaterial(id uint) error
	AddMaterialToItem(itemID, materialID uint, quantity int) error
	RemoveMaterialFromItem(itemID, materialID uint) error
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
	err := config.DB.Preload("Materials.Material").First(&item, id).Error
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
	err = query.Preload("Materials.Material").Offset(offset).Limit(pageSize).Find(&items).Error

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

// CreateMaterial 创建材料记录
func (r *examinationItemRepository) CreateMaterial(material *models.Material) error {
	return config.DB.Create(material).Error
}

// FindMaterialByID 通过ID查找材料记录
func (r *examinationItemRepository) FindMaterialByID(id uint) (*models.Material, error) {
	var material models.Material
	err := config.DB.First(&material, id).Error
	return &material, err
}

// FindAllMaterials 查找所有材料记录
func (r *examinationItemRepository) FindAllMaterials() ([]models.Material, error) {
	var materials []models.Material
	err := config.DB.Find(&materials).Error
	return materials, err
}

// UpdateMaterial 更新材料记录
func (r *examinationItemRepository) UpdateMaterial(material *models.Material) error {
	return config.DB.Save(material).Error
}

// DeleteMaterial 删除材料记录
func (r *examinationItemRepository) DeleteMaterial(id uint) error {
	return config.DB.Delete(&models.Material{}, id).Error
}

// AddMaterialToItem 向检查项目添加材料
func (r *examinationItemRepository) AddMaterialToItem(itemID, materialID uint, quantity int) error {
	itemMaterial := models.ExaminationItemMaterial{
		ExaminationItemID: itemID,
		MaterialID:        materialID,
		Quantity:          quantity,
	}
	return config.DB.Create(&itemMaterial).Error
}

// RemoveMaterialFromItem 从检查项目中移除材料
func (r *examinationItemRepository) RemoveMaterialFromItem(itemID, materialID uint) error {
	return config.DB.Where("examination_item_id = ? AND material_id = ?", itemID, materialID).Delete(&models.ExaminationItemMaterial{}).Error
}
