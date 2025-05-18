package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/fe/models"
	"github.com/yourusername/fe/services"
)

// DoctorController 医生信息表控制器
type DoctorController struct {
	doctorService services.DoctorService
}

// NewDoctorController 创建医生信息表控制器
func NewDoctorController(doctorService services.DoctorService) *DoctorController {
	return &DoctorController{
		doctorService: doctorService,
	}
}

// Create 创建医生记录
func (c *DoctorController) Create(ctx *gin.Context) {
	var doctor models.Doctor
	if err := ctx.ShouldBindJSON(&doctor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.doctorService.CreateDoctor(&doctor); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, doctor)
}

// GetByID 获取医生记录详情
func (c *DoctorController) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	doctor, err := c.doctorService.GetDoctorByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "未找到记录"})
		return
	}

	ctx.JSON(http.StatusOK, doctor)
}

// GetAll 获取医生记录列表
func (c *DoctorController) GetAll(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))

	params := make(map[string]interface{})

	if name := ctx.Query("name"); name != "" {
		params["name"] = name
	}
	if department := ctx.Query("department"); department != "" {
		params["department"] = department
	}

	doctors, total, err := c.doctorService.GetAllDoctors(params, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  doctors,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}

// Update 更新医生记录
func (c *DoctorController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	doctor, err := c.doctorService.GetDoctorByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "未找到记录"})
		return
	}

	if err := ctx.ShouldBindJSON(doctor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.doctorService.UpdateDoctor(doctor); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, doctor)
}

// Delete 删除医生记录
func (c *DoctorController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	if err := c.doctorService.DeleteDoctor(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "记录删除成功"})
}
