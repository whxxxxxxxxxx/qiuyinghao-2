package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/fe/controllers"
	"github.com/yourusername/fe/middlewares"
	"github.com/yourusername/fe/repositories"
	"github.com/yourusername/fe/services"
)

// SetupUserRoutes 设置用户相关路由
func SetupUserRoutes(router *gin.Engine) {
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	userGroup := router.Group("/api/users")
	{
		userGroup.GET("/", userController.GetAll)
		userGroup.GET("/:id", userController.GetByID)
		userGroup.POST("/", userController.Create)

		// 需要认证的路由
		authorized := userGroup.Group("/")
		authorized.Use(middlewares.AuthMiddleware())
		{
			authorized.PUT("/:id", userController.Update)
			authorized.DELETE("/:id", userController.Delete)
		}
	}
}
