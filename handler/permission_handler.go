package handler

import (
	"github.com/gin-gonic/gin"
)

func VerifyIdentity(context *gin.Context) {
	// 通过用户标识匹配密码凭证
	// 返回验证结果
}

func VerifyPermission(context *gin.Context) {
	// pipeline 校验用户权限
}