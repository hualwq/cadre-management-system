package app

import (
	"cadre-management/pkg/e"
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	// 尝试绑定请求参数（根据 Content-Type 自动选择 BindJSON、BindForm 等）
	err := c.Bind(form)
	if err != nil {
		// 打印绑定错误信息
		fmt.Printf("参数绑定失败: %v\n", err)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	// 进行字段验证（github.com/astaxie/beego/validation）
	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		fmt.Printf("验证处理异常: %v\n", err)
		return http.StatusInternalServerError, e.ERROR
	}
	if !check {
		for _, verr := range valid.Errors {
			fmt.Printf("验证失败 - 字段: %s, 错误: %s\n", verr.Field, verr.Message)
		}
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}
