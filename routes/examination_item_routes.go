package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/fe/controllers"
	"github.com/yourusername/fe/repositories"
	"github.com/yourusername/fe/services"
)

// SetupExaminationItemRoutes 设置检查项目表相关路由
func SetupExaminationItemRoutes(router *gin.Engine) {
	itemRepo := repositories.NewExaminationItemRepository()
	itemService := services.NewExaminationItemService(itemRepo)
	itemController := controllers.NewExaminationItemController(itemService)

	itemGroup := router.Group("/api/examination-items")
	{
		itemGroup.POST("", itemController.Create)
		itemGroup.GET("", itemController.GetAll)
		itemGroup.GET("/:id", itemController.GetByID)
		itemGroup.PUT("/:id", itemController.Update)
		itemGroup.DELETE("/:id", itemController.Delete)
	}
}
