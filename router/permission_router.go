package router

import (
	"PasteMeGoAccount/handler"
	"github.com/gin-gonic/gin"
)

func LoadPermissionRouter(ginEngine *gin.Engine) {
	routerGroup := ginEngine.Group("permission")
	{
		// 身份校验
		routerGroup.GET("identities", handler.VerifyIdentity)
		// 权限校验
		routerGroup.GET("permissions", handler.VerifyPermission)
	}
}
