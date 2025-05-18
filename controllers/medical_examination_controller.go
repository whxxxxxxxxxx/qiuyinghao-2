package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"github.com/yourusername/fe/models"
	"github.com/yourusername/fe/services"
)

// MedicalExaminationController 病历检查表控制器
type MedicalExaminationController struct {
	examService services.MedicalExaminationService
}

// NewMedicalExaminationController 创建病历检查表控制器
func NewMedicalExaminationController(examService services.MedicalExaminationService) *MedicalExaminationController {
	return &MedicalExaminationController{
		examService: examService,
	}
}

// Create 创建病历检查记录
func (c *MedicalExaminationController) Create(ctx *gin.Context) {
	var exam models.MedicalExamination
	if err := ctx.ShouldBindJSON(&exam); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.examService.CreateExamination(&exam); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, exam)
}

// GetByID 获取病历检查记录详情
func (c *MedicalExaminationController) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	exam, err := c.examService.GetExaminationByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "未找到记录"})
		return
	}

	ctx.JSON(http.StatusOK, exam)
}

// GetAll 获取病历检查记录列表（支持筛选和分页）
func (c *MedicalExaminationController) GetAll(ctx *gin.Context) {
	// 获取分页参数
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))

	// 获取筛选参数
	params := make(map[string]interface{})

	if patientName := ctx.Query("patientName"); patientName != "" {
		params["patientName"] = patientName
	}
	if examinationItem := ctx.Query("examinationItem"); examinationItem != "" {
		//转化为int类型
		examinationItemID, err := strconv.ParseUint(examinationItem, 10, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的检查项目ID"})
			return
		}
		//将字符串转化为int类型
		params["examinationItem"] = uint(examinationItemID)

	}
	if department := ctx.Query("department"); department != "" {

		params["department"] = department
	}
	if doctor := ctx.Query("doctor"); doctor != "" {
		//转化为int类型
		doctorID, err := strconv.ParseUint(doctor, 10, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的医生ID"})
			return
		}
		params["doctor"] = uint(doctorID)
	}
	// 其他筛选参数可以根据需要添加

	exams, total, err := c.examService.GetAllExaminations(params, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  exams,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}

// Update 更新病历检查记录
func (c *MedicalExaminationController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	exam, err := c.examService.GetExaminationByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "未找到记录"})
		return
	}

	if err := ctx.ShouldBindJSON(exam); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.examService.UpdateExamination(exam); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exam)
}

// Delete 删除病历检查记录
func (c *MedicalExaminationController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	if err := c.examService.DeleteExamination(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "记录删除成功"})
}

// ExportToExcel 导出病历检查记录到Excel
func (c *MedicalExaminationController) ExportToExcel(ctx *gin.Context) {
	// 获取筛选参数
	params := make(map[string]interface{})

	if patientName := ctx.Query("patientName"); patientName != "" {
		params["patientName"] = patientName
	}
	if doctorIDStr := ctx.Query("doctorID"); doctorIDStr != "" {
		if doctorID, err := strconv.ParseUint(doctorIDStr, 10, 32); err == nil {
			params["doctorID"] = uint(doctorID)
		}
	}
	if itemIDStr := ctx.Query("examinationItemID"); itemIDStr != "" {
		if itemID, err := strconv.ParseUint(itemIDStr, 10, 32); err == nil {
			params["examinationItemID"] = uint(itemID)
		}
	}

	// 获取所有符合条件的记录
	exams, err := c.examService.ExportExaminations(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "获取数据失败: " + err.Error()})
		return
	}

	// 创建Excel文件
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// 创建工作表
	sheetName := "病历检查记录"
	_, err = f.NewSheet(sheetName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "创建Excel工作表失败: " + err.Error()})
		return
	}

	// 设置表头
	headers := []string{"ID", "病号", "检查项目", "检查次数", "总金额", "开单医生", "成本分配率", "创建时间"}
	for i, header := range headers {
		cell := fmt.Sprintf("%c%d", 'A'+i, 1)
		f.SetCellValue(sheetName, cell, header)
	}

	// 填充数据
	for i, exam := range exams {
		row := i + 2 // 从第2行开始（第1行是表头）
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), exam.ID)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), exam.PatientName)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), exam.ExaminationItem.Name)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), exam.ExaminationCount)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), exam.TotalAmount)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), exam.Doctor.Name)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), exam.CostAllocationRate)
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), exam.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	// 设置列宽
	f.SetColWidth(sheetName, "A", "H", 15)

	// 设置文件名
	filename := fmt.Sprintf("病历检查记录_%s.xlsx", time.Now().Format("20060102150405"))

	// 设置响应头
	ctx.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	// 将Excel文件写入响应
	if err := f.Write(ctx.Writer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "导出Excel失败: " + err.Error()})
		return
	}
}
