package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/fe/config"
	"github.com/yourusername/fe/models"
)

type DashboardController struct{}

func NewDashboardController() *DashboardController {
	return &DashboardController{}
}

func (c *DashboardController) Summary(ctx *gin.Context) {
	var doctorCount int64
	var departmentCount int64
	var todayVisitCount int64
	var monthVisitCount int64
	var todayIncome float64
	var monthIncome float64
	var todayMaterialCost float64
	var monthMaterialCost float64

	today := time.Now().Format("2006-01-02")
	month := time.Now().Format("2006-01")

	// 医生数量
	config.DB.Model(&models.Doctor{}).Count(&doctorCount)
	// 科室数量
	config.DB.Model(&models.Department{}).Count(&departmentCount)
	// 今日门诊量
	config.DB.Model(&models.MedicalExamination{}).Where("DATE(created_at) = ?", today).Count(&todayVisitCount)
	// 本月门诊量
	config.DB.Model(&models.MedicalExamination{}).Where("DATE_FORMAT(created_at, '%Y-%m') = ?", month).Count(&monthVisitCount)
	// 今日收入
	config.DB.Model(&models.MedicalExamination{}).Where("DATE(created_at) = ?", today).Select("COALESCE(SUM(total_amount),0)").Scan(&todayIncome)
	// 本月收入
	config.DB.Model(&models.MedicalExamination{}).Where("DATE_FORMAT(created_at, '%Y-%m') = ?", month).Select("COALESCE(SUM(total_amount),0)").Scan(&monthIncome)

	// 材料消耗（遍历今日/本月所有检查项目，累加材料消耗）
	var todayExams []models.MedicalExamination
	config.DB.Preload("ExaminationItem.Materials.Material").Where("DATE(created_at) = ?", today).Find(&todayExams)
	for _, exam := range todayExams {
		for _, m := range exam.ExaminationItem.Materials {
			todayMaterialCost += float64(m.Quantity) * m.Material.Price
		}
	}
	var monthExams []models.MedicalExamination
	config.DB.Preload("ExaminationItem.Materials.Material").Where("DATE_FORMAT(created_at, '%Y-%m') = ?", month).Find(&monthExams)
	for _, exam := range monthExams {
		for _, m := range exam.ExaminationItem.Materials {
			monthMaterialCost += float64(m.Quantity) * m.Material.Price
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"doctorCount":       doctorCount,
		"departmentCount":   departmentCount,
		"todayVisitCount":   todayVisitCount,
		"monthVisitCount":   monthVisitCount,
		"todayIncome":       todayIncome,
		"monthIncome":       monthIncome,
		"todayMaterialCost": todayMaterialCost,
		"monthMaterialCost": monthMaterialCost,
	})
}
