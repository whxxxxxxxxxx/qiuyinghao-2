package models

import (
	"gorm.io/gorm"
)

// ExaminationItem 检查项目表模型
type ExaminationItem struct {
	gorm.Model
	Name            string  `json:"name" gorm:"type:varchar(255);not null;comment:项目名称"`
	Amount          float64 `json:"amount" gorm:"not null;comment:金额"`
	InsuranceAmount float64 `json:"insuranceAmount" gorm:"not null;comment:医保结算金额"`
	CostRatio       float64 `json:"costRatio" gorm:"not null;default:0.5;comment:成本比例"`
	DepartmentRatio float64 `json:"departmentRatio" gorm:"not null;default:0.5;comment:科室分配比例"`
}

// CalculateCost 计算成本
func (e *ExaminationItem) CalculateCost() float64 {
	return e.Amount * e.CostRatio
}

// CalculateDepartmentAmount 计算分配给开单科室的金额
func (e *ExaminationItem) CalculateDepartmentAmount() float64 {
	return e.CalculateCost() * e.DepartmentRatio
}
