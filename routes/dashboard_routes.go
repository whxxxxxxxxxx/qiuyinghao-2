package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/fe/controllers"
)

func SetupDashboardRoutes(router *gin.Engine) {
	controller := controllers.NewDashboardController()
	group := router.Group("/api/dashboard")
	{
		group.GET("/summary", controller.Summary)
	}
}
