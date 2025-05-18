package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/fe/controllers"
	"github.com/yourusername/fe/repositories"
	"github.com/yourusername/fe/services"
)

// SetupDoctorRoutes 设置医生信息表相关路由
func SetupDoctorRoutes(router *gin.Engine) {
	doctorRepo := repositories.NewDoctorRepository()
	doctorService := services.NewDoctorService(doctorRepo)
	doctorController := controllers.NewDoctorController(doctorService)

	doctorGroup := router.Group("/api/doctors")
	{
		doctorGroup.POST("", doctorController.Create)
		doctorGroup.GET("", doctorController.GetAll)
		doctorGroup.GET("/:id", doctorController.GetByID)
		doctorGroup.PUT("/:id", doctorController.Update)
		doctorGroup.DELETE("/:id", doctorController.Delete)
	}
}
