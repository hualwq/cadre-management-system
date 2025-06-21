package v1

import (
	"cadre-management/models"
	"cadre-management/pkg/app"
	"cadre-management/pkg/e"
	"cadre-management/services/Sys_admin"
	"net/http"
	"strconv"

	"github.com/unknwon/com"

	"github.com/gin-gonic/gin"
)

type ChangeUserROleForm struct {
	CadreID string `json:"user_id"`
	Role    string `json:"role"`
}

func GetUserByPage(c *gin.Context) {
	appG := app.Gin{C: c}
	pageNumstr := c.Query("page")
	pageSizestr := c.Query("pagesize")
	if pageNumstr == "" || pageSizestr == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	pageNum := com.StrTo(pageNumstr).MustInt()
	pageSize := com.StrTo(pageSizestr).MustInt()

	userService := Sys_admin.GetUser{
		PageNum:  pageNum,
		PageSize: pageSize,
	}

	users, err := userService.GetUserByPage(pageNum, pageSize)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, users)
}

func GetAllUser(c *gin.Context) {
	appG := app.Gin{C: c}

	// 创建service实例
	userService := Sys_admin.User{}

	// 调用service层方法
	users, err := userService.GetAllUser()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_USERLIST_FAIL, nil)
		return
	}

	// 返回成功响应
	appG.Response(http.StatusOK, e.SUCCESS, users)
}

func ChangeUserRole(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form ChangeUserROleForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	// 调用服务层方法
	userService := Sys_admin.ChangeUserRole{
		CadreID: form.CadreID,
		Role:    form.Role,
	}
	err := userService.ChangeUserRole(userService.CadreID, userService.Role)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 响应成功
	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"message": "用户角色更改成功",
	})
}

// CreateDepartment 创建院系
func CreateDepartment(c *gin.Context) {
	appG := app.Gin{C: c}

	var department models.Department
	if err := c.ShouldBindJSON(&department); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	if err := models.CreateDepartment(&department); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, gin.H{
			"error": err.Error(),
		})
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"message":    "院系创建成功",
		"department": department,
	})
}

// UpdateDepartment 更新院系
func UpdateDepartment(c *gin.Context) {
	appG := app.Gin{C: c}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	var department models.Department
	if err := c.ShouldBindJSON(&department); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	if err := models.UpdateDepartment(uint(id), &department); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, gin.H{
			"error": err.Error(),
		})
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"message": "院系更新成功",
	})
}

// DeleteDepartment 删除院系
func DeleteDepartment(c *gin.Context) {
	appG := app.Gin{C: c}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	if err := models.DeleteDepartment(uint(id)); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, gin.H{
			"error": err.Error(),
		})
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"message": "院系删除成功",
	})
}

// GetDepartmentByID 根据ID获取院系
func GetDepartmentByID(c *gin.Context) {
	appG := app.Gin{C: c}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	department, err := models.GetDepartmentByID(uint(id))
	if err != nil {
		appG.Response(http.StatusNotFound, e.ERROR, gin.H{
			"error": "院系不存在",
		})
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		"department": department,
	})
}

type ChangeUserRoleForm struct {
	Role         string `json:"role" binding:"required"`
	DepartmentID int    `json:"department_id"`
}

// PUT /sysadmin/user/:id/role
func ChangeUserRoleByID(c *gin.Context) {
	appG := app.Gin{C: c}
	userID := c.Param("id")
	var form ChangeUserRoleForm
	if err := c.ShouldBindJSON(&form); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	changeUserRole := Sys_admin.ChangeUserRole{
		CadreID: userID,
		Role:    form.Role,
	}
	err := changeUserRole.ChangeUserRoleByID(changeUserRole.CadreID, changeUserRole.Role)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// GET /sysadmin/user/role?role=xxx
func GetUserRoleList(c *gin.Context) {
	appG := app.Gin{C: c}
	role := c.Query("role")
	if role == "" {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	users, err := models.GetUsersByRole(role)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, gin.H{"users": users})
}
