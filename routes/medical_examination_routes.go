package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/fe/controllers"
	"github.com/yourusername/fe/repositories"
	"github.com/yourusername/fe/services"
)

// SetupMedicalExaminationRoutes 设置病历检查表相关路由
func SetupMedicalExaminationRoutes(router *gin.Engine) {
	examRepo := repositories.NewMedicalExaminationRepository()
	examService := services.NewMedicalExaminationService(examRepo)
	examController := controllers.NewMedicalExaminationController(examService)

	examGroup := router.Group("/api/medical-examinations")
	{
		examGroup.POST("", examController.Create)
		examGroup.GET("", examController.GetAll)
		examGroup.GET("/:id", examController.GetByID)
		examGroup.PUT("/:id", examController.Update)
		examGroup.DELETE("/:id", examController.Delete)
		examGroup.GET("/export", examController.ExportToExcel) // 添加导出Excel的路由
	}
}
