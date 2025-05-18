package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yourusername/fe/config"
	"github.com/yourusername/fe/models"
	"github.com/yourusername/fe/routes"
)

func main() {
	// 加载配置
	appConfig := config.LoadConfig()

	// 连接数据库
	config.ConnectDatabase(appConfig)

	// 自动迁移数据库表结构
	err := config.DB.AutoMigrate(&models.User{}, &models.Doctor{}, &models.MedicalExamination{}, &models.ExaminationItem{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	// 创建Gin实例
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // 您的前端应用地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 设置路由
	routes.SetupRoutes(router)

	// 启动服务器
	log.Printf("Server running on port %s", appConfig.ServerPort)
	if err := router.Run(":" + appConfig.ServerPort); err != nil {
		log.Fatal("Failed to start server: ", err)
	}

}
