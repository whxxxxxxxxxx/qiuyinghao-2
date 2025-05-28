package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置所有路由
func SetupRoutes(router *gin.Engine) {
	// 设置用户路由
	SetupUserRoutes(router)

	// 设置医生路由
	SetupDoctorRoutes(router)

	// 设置检查项目路由
	SetupExaminationItemRoutes(router)

	// 设置病历检查路由
	SetupMedicalExaminationRoutes(router)

	// 设置测试路由
	SetupTestRoutes(router)

	// 设置科室路由
	SetupDepartmentRoutes(router)

	// 设置医院运营看板路由
	SetupDashboardRoutes(router)
}
