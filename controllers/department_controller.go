package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/fe/models"
	"github.com/yourusername/fe/services"
)

type DepartmentController struct {
	service services.DepartmentService
}

func NewDepartmentController(service services.DepartmentService) *DepartmentController {
	return &DepartmentController{service: service}
}

func (c *DepartmentController) Create(ctx *gin.Context) {
	var dept models.Department
	if err := ctx.ShouldBindJSON(&dept); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.CreateDepartment(&dept); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, dept)
}

func (c *DepartmentController) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}
	dept, err := c.service.GetDepartmentByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "未找到科室"})
		return
	}
	ctx.JSON(http.StatusOK, dept)
}

func (c *DepartmentController) GetAll(ctx *gin.Context) {
	depts, err := c.service.GetAllDepartments()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": depts, "total": len(depts)})
}

func (c *DepartmentController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}
	var dept models.Department
	if err := ctx.ShouldBindJSON(&dept); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dept.ID = uint(id)
	if err := c.service.UpdateDepartment(&dept); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, dept)
}

func (c *DepartmentController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}
	if err := c.service.DeleteDepartment(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "科室删除成功"})
}
