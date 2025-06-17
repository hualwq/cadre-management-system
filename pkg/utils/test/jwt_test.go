package test

import (
	"cadre-management/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAndParseToken(t *testing.T) {
	// 准备测试数据
	userID := "111111111"
	password := "123456"
	role := "cadre"

	// 生成 JWT
	token, err := utils.GenerateToken(userID, password, role)
	assert.NoError(t, err)

	// 解析 JWT
	claims, err := utils.ParseToken(token)
	assert.NoError(t, err)

	// 比较 user_id
	assert.Equal(t, userID, claims.UserID)
}
