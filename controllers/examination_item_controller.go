package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/fe/models"
	"github.com/yourusername/fe/services"
)

// ExaminationItemController 检查项目表控制器
type ExaminationItemController struct {
	itemService services.ExaminationItemService
}

// NewExaminationItemController 创建检查项目表控制器
func NewExaminationItemController(itemService services.ExaminationItemService) *ExaminationItemController {
	return &ExaminationItemController{
		itemService: itemService,
	}
}

// Create 创建检查项目记录
func (c *ExaminationItemController) Create(ctx *gin.Context) {
	var item models.ExaminationItem
	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.itemService.CreateExaminationItem(&item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, item)
}

// GetByID 获取检查项目记录详情
func (c *ExaminationItemController) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	item, err := c.itemService.GetExaminationItemByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "未找到记录"})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

// GetAll 获取检查项目记录列表
func (c *ExaminationItemController) GetAll(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))

	params := make(map[string]interface{})

	if name := ctx.Query("name"); name != "" {
		params["name"] = name
	}

	items, total, err := c.itemService.GetAllExaminationItems(params, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  items,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}

// Update 更新检查项目记录
func (c *ExaminationItemController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	item, err := c.itemService.GetExaminationItemByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "未找到记录"})
		return
	}

	if err := ctx.ShouldBindJSON(item); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.itemService.UpdateExaminationItem(item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

// Delete 删除检查项目记录
func (c *ExaminationItemController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	if err := c.itemService.DeleteExaminationItem(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "记录删除成功"})
}
