package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/fe/controllers"
)

// SetupTestRoutes 设置测试相关路由
func SetupTestRoutes(router *gin.Engine) {
	testController := controllers.NewTestController()

	// 测试路由组
	testGroup := router.Group("/api/test")
	{
		testGroup.GET("/ping", testController.Ping)
	}
}
