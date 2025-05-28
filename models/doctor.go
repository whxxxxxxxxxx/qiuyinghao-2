package models

import (
	"gorm.io/gorm"
)

// Doctor 医生信息表模型
type Doctor struct {
	gorm.Model
	Name         string     `json:"name" gorm:"type:varchar(100);not null;comment:医生名"`
	DepartmentID uint       `json:"departmentId" gorm:"not null;comment:科室ID"`
	Department   Department `json:"department" gorm:"foreignKey:DepartmentID"`
}

// Department 科室表模型
type Department struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(100);not null;uniqueIndex;comment:科室名称"`
	Description string `json:"description" gorm:"type:text;comment:科室介绍"`
}
