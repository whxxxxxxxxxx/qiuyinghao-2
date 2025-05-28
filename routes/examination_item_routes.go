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

	// 材料管理相关路由
	materialGroup := router.Group("/api/materials")
	{
		materialGroup.POST("", itemController.CreateMaterial)
		materialGroup.GET("", itemController.GetAllMaterials)
		materialGroup.GET("/:id", itemController.GetMaterialByID)
		materialGroup.PUT("/:id", itemController.UpdateMaterial)
		materialGroup.DELETE("/:id", itemController.DeleteMaterial)
	}

	// 检查项目材料关联路由
	itemMaterialGroup := router.Group("/api/examination-items/materials/:itemId")
	{
		itemMaterialGroup.POST("/:materialId", itemController.AddMaterialToItem)
		itemMaterialGroup.DELETE("/:materialId", itemController.RemoveMaterialFromItem)
	}

	// 检查项目基本路由
	itemGroup := router.Group("/api/examination-items")
	{
		itemGroup.POST("", itemController.Create)
		itemGroup.GET("", itemController.GetAll)
		itemGroup.GET("/:id", itemController.GetByID)
		itemGroup.PUT("/:id", itemController.Update)
		itemGroup.DELETE("/:id", itemController.Delete)
	}
}
