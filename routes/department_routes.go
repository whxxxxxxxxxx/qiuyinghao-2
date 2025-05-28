package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/fe/controllers"
	"github.com/yourusername/fe/repositories"
	"github.com/yourusername/fe/services"
)

func SetupDepartmentRoutes(router *gin.Engine) {
	repo := repositories.NewDepartmentRepository()
	service := services.NewDepartmentService(repo)
	controller := controllers.NewDepartmentController(service)

	group := router.Group("/api/departments")
	{
		group.POST("", controller.Create)
		group.GET("", controller.GetAll)
		group.GET(":id", controller.GetByID)
		group.PUT(":id", controller.Update)
		group.DELETE(":id", controller.Delete)
	}
}
