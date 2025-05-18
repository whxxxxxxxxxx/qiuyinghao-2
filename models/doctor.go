package models

import (
	"gorm.io/gorm"
)

// Doctor 医生信息表模型
type Doctor struct {
	gorm.Model
	Name       string `json:"name" gorm:"type:varchar(100);not null;comment:医生名"`
	Department string `json:"department" gorm:"type:varchar(100);not null;comment:科室"`
}
