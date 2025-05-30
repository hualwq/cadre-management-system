package v1

import (
	"cadre-management/pkg/app"
	"cadre-management/pkg/e"
	"cadre-management/services/sys_admin"
	"net/http"

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

	userService := sys_admin.GetUser{
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
	userService := sys_admin.User{}

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
	userService := sys_admin.ChangeUserRole{
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
