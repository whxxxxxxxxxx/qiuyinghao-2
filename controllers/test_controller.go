package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// TestController 测试控制器
type TestController struct{}

// NewTestController 创建测试控制器
func NewTestController() *TestController {
	return &TestController{}
}

// Ping 测试接口
func (c *TestController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":   "pong",
		"timestamp": time.Now().Format(time.RFC3339),
		"status":    "success",
	})
}
