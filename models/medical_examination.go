package models

import (
	"gorm.io/gorm"
)

// MedicalExamination 病历检查表模型
type MedicalExamination struct {
	gorm.Model
	PatientName        string          `json:"patientName" gorm:"type:varchar(100);not null;comment:病号"`
	ExaminationItemID  uint            `json:"examinationItemID" gorm:"not null;comment:检查项目ID"`
	ExaminationItem    ExaminationItem `json:"examinationItem" gorm:"foreignKey:ExaminationItemID"`
	ExaminationCount   int             `json:"examinationCount" gorm:"not null;default:1;comment:检查次数"`
	TotalAmount        float64         `json:"totalAmount" gorm:"not null;comment:总金额"`
	DoctorID           uint            `json:"doctorID" gorm:"not null;comment:开单医生ID"`
	Doctor             Doctor          `json:"doctor" gorm:"foreignKey:DoctorID"`
	CostAllocationRate float64         `json:"costAllocationRate" gorm:"not null;comment:成本分配率"`
}
