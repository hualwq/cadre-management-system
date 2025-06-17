package e

var MsgFlags = map[int]string{
	SUCCESS:                            "ok",
	ERROR:                              "fail",
	INVALID_PARAMS:                     "请求参数错误",
	ERROR_PAGE:                         "获取评估列表失败",
	ERROR_PAGE_PARAMS:                  "分页参数错误",
	ERROR_ADD_ASSESSMENT_FAIL:          "添加评价失败",
	ERROR_GET_CADREINFO_FAIL:           "获取所有标签失败",
	ERROR_EDIT_POSITIONHISTORY_FAIL:    "编辑干部经历失败",
	ERROR_ADD_CADRE_FAIL:               "新增干部信息失败",
	ERROR_GET_ROLE:                     "获取用户权限失败",
	ERROR_PERMISSION:                   "权限检查错误",
	ERROR_ADD_POSITIONHISTORY_FAIL:     "添加任职信息失败",
	ERROR_CONFIRM_POSITIONHISTORY_FAIL: "确认提交任职信息失败",
	ERROR_NOT_EXIST_ARTICLE:            "该文章不存在",
	ERROR_ADD_ARTICLE_FAIL:             "新增文章失败",
	ERROR_AUTH_CHECK_FAIL:              "用户权限检查错误",
	ERROR_ADD_POSITION_FAIL:            "添加工作经历失败",
	ERROR_CHANGE_ROLE_FAIL:             "修改用户权限失败",
	ERROR_COUNT_ARTICLE_FAIL:           "统计文章失败",
	ERROR_DATABASE:                     "数据库错误",
	ERROR_USER_NOT_LOGIN:               "用户未登录",
	ERROR_GET_USERLIST_FAIL:            "获取用户列表失败",
	ERROR_USER_CHECK_TOKEN_FAIL:        "Token鉴权失败",
	ERROR_USER_CHECK_TOKEN_TIMEOUT:     "Token已超时",
	ERROR_AUTH_TOKEN:                   "Token生成失败",
	ERROR_AUTH:                         "用户名密码验证错误",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:       "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:      "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT:    "校验图片错误，图片格式或大小有问题",
	ERROR_NOT_EXSIT_FAMILYMEMBER:       "没有家庭成员",
	ERROR_NOT_EXSIT_CADREINFO:          "干部信息不存在",
	ERROR_EDIT_CADRE_FAIL:              "编辑干部信息失败",
	ERROR_ADD_RESUME_FAIL:              "新增简历信息失败",
	ERROR_ADD_FAMILYMEBER:              "新增家庭成员失败",
	ERROR_NOT_EXSIT_RESUME:             "不存在简历信息",
	ERROR_GET_POSITION_HISTORIES_FAIL:  "获取任职信息失败",
	ERROR_GET_ASSESSMENTS_FAIL:         "获取待审核的考核信息失败",
	ERROR_CHECK_EXIST_ASSESEMENT_FAIL:  "检查存在考核信息失败",
	ERROR_NOT_EXIST_ASSESEMENT:         "不存在该考核信息",
	ERROR_FAMILYMEMBERIES_FAIL:         "不存在相关家庭成员",
	ERROR_RESUME_FAIL:                  "查找简历失败",
	ERROR_EDIT_FAMILYMEMBER_FAIL:       "编辑家庭成员失败",
	ERROR_EDIT_RESUME_FAIL:             "编辑简历失败",
	ERROR_NOT_EXSIT_POSITIONHISTORY:    "不存在工作经历信息",
	ERROR_DELETE_POSITIONHISTORY_FAIL:  "删除工作经历失败",
	ERROR_DELETE_CADRE_FAIL:            "删除干部信息失败",
	ERROR_GET_POEXPMOD_FAIL:            "获取年度经历失败",
	ERROR_DELETE_POSEXP_FAIL:           "删除年度经历失败",
	ERROR_DELETE_ASSESSMENT_FAIL:       "删除考核失败",
	ERROR_NOT_EXSIT_ASSESSMENT:         "不存在考核信息",
	ERROR_EDIT_ASSESSMENT_FAIL:         "编辑考核信息失败",
	ERROR_USER_NOT_LOGGED_IN:           "请重新登录",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
