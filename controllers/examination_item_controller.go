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

// CreateMaterial 创建材料记录
func (c *ExaminationItemController) CreateMaterial(ctx *gin.Context) {
	var material models.Material
	if err := ctx.ShouldBindJSON(&material); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.itemService.CreateMaterial(&material); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, material)
}

// GetMaterialByID 获取材料记录详情
func (c *ExaminationItemController) GetMaterialByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	material, err := c.itemService.GetMaterialByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "未找到记录"})
		return
	}

	ctx.JSON(http.StatusOK, material)
}

// GetAllMaterials 获取材料记录列表
func (c *ExaminationItemController) GetAllMaterials(ctx *gin.Context) {
	materials, err := c.itemService.GetAllMaterials()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, materials)
}

// UpdateMaterial 更新材料记录
func (c *ExaminationItemController) UpdateMaterial(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	material, err := c.itemService.GetMaterialByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "未找到记录"})
		return
	}

	if err := ctx.ShouldBindJSON(material); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.itemService.UpdateMaterial(material); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, material)
}

// DeleteMaterial 删除材料记录
func (c *ExaminationItemController) DeleteMaterial(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	if err := c.itemService.DeleteMaterial(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "记录删除成功"})
}

// AddMaterialToItem 向检查项目添加材料
func (c *ExaminationItemController) AddMaterialToItem(ctx *gin.Context) {
	itemID, err := strconv.ParseUint(ctx.Param("itemId"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的项目ID"})
		return
	}

	materialID, err := strconv.ParseUint(ctx.Param("materialId"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的材料ID"})
		return
	}

	var req struct {
		Quantity int `json:"quantity"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil || req.Quantity <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的数量"})
		return
	}
	quantity := req.Quantity

	if err := c.itemService.AddMaterialToItem(uint(itemID), uint(materialID), quantity); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "材料添加成功"})
}

// RemoveMaterialFromItem 从检查项目中移除材料
func (c *ExaminationItemController) RemoveMaterialFromItem(ctx *gin.Context) {
	itemID, err := strconv.ParseUint(ctx.Param("itemId"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的项目ID"})
		return
	}

	materialID, err := strconv.ParseUint(ctx.Param("materialId"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的材料ID"})
		return
	}

	if err := c.itemService.RemoveMaterialFromItem(uint(itemID), uint(materialID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "材料移除成功"})
}
