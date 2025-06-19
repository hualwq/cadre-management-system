package v1

import (
	"cadre-management/models"
	"cadre-management/services/Department_service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var deptService *Department_service.DepartmentService

func InitDepartmentService() {
	deptService = &Department_service.DepartmentService{DB: models.GetDB()}
}

// 新增院系
func CreateDepartment(c *gin.Context) {
	type Req struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := deptService.CreateDepartment(req.Name, req.Description); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

// 删除院系
func DeleteDepartment(c *gin.Context) {
	id := c.Param("id")
	var deptID uint
	_, err := fmt.Sscanf(id, "%d", &deptID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := deptService.DeleteDepartment(deptID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

// 修改院系
func UpdateDepartment(c *gin.Context) {
	type Req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	id := c.Param("id")
	var deptID uint
	_, err := fmt.Sscanf(id, "%d", &deptID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := deptService.UpdateDepartment(deptID, req.Name, req.Description); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

// 查询所有院系
func ListDepartments(c *gin.Context) {
	depts, err := deptService.ListDepartments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": depts})
}

// 设置院系管理员
func SetDepartmentAdmin(c *gin.Context) {
	type Req struct {
		UserID       string `json:"user_id" binding:"required"`
		DepartmentID uint   `json:"department_id" binding:"required"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := deptService.SetDepartmentAdmin(req.UserID, req.DepartmentID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

// 取消院系管理员
func UnsetDepartmentAdmin(c *gin.Context) {
	type Req struct {
		UserID       string `json:"user_id" binding:"required"`
		DepartmentID uint   `json:"department_id" binding:"required"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := deptService.UnsetDepartmentAdmin(req.UserID, req.DepartmentID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

// 查询院系管理员
func GetDepartmentAdmins(c *gin.Context) {
	id := c.Query("department_id")
	var deptID uint
	_, err := fmt.Sscanf(id, "%d", &deptID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	users, err := deptService.GetDepartmentAdmins(deptID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}
