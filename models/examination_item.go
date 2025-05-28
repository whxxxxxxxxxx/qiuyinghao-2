package models

import (
	"gorm.io/gorm"
)

// Material 材料表模型
type Material struct {
	gorm.Model
	Name  string  `json:"name" gorm:"type:varchar(255);not null;comment:材料名称"`
	Price float64 `json:"price" gorm:"not null;comment:材料价格"`
}

// ExaminationItemMaterial 项目材料关联表模型
type ExaminationItemMaterial struct {
	gorm.Model
	ExaminationItemID uint     `json:"examinationItemId" gorm:"not null;comment:检查项目ID"`
	MaterialID        uint     `json:"materialId" gorm:"not null;comment:材料ID"`
	Quantity          int      `json:"quantity" gorm:"not null;comment:材料数量"`
	Material          Material `json:"material" gorm:"foreignKey:MaterialID"`
}

// ExaminationItem 检查项目表模型
type ExaminationItem struct {
	gorm.Model
	Name            string                    `json:"name" gorm:"type:varchar(255);not null;comment:项目名称"`
	Amount          float64                   `json:"amount" gorm:"not null;comment:金额"`
	InsuranceAmount float64                   `json:"insuranceAmount" gorm:"not null;comment:医保结算金额"`
	CostRatio       float64                   `json:"costRatio" gorm:"not null;default:0.5;comment:成本比例"`
	DepartmentRatio float64                   `json:"departmentRatio" gorm:"not null;default:0.5;comment:科室分配比例"`
	Materials       []ExaminationItemMaterial `json:"materials" gorm:"foreignKey:ExaminationItemID"`
}

// CalculateCost 计算成本
func (e *ExaminationItem) CalculateCost() float64 {
	return e.Amount * e.CostRatio
}

// CalculateDepartmentAmount 计算分配给开单科室的金额
func (e *ExaminationItem) CalculateDepartmentAmount() float64 {
	return e.CalculateCost() * e.DepartmentRatio
}
