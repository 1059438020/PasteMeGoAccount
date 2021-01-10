package router

import (
	"PasteMeGoAccount/handler"
	"github.com/gin-gonic/gin"
)

func LoadAccountRouter(ginEngine *gin.Engine) {
	routerGroup := ginEngine.Group("account")
	{
		// 注册账号
		routerGroup.POST("accounts", handler.RegisterAccount)
		// 冻结账号
		routerGroup.PUT("accounts", handler.FrozenAccount)
		// 修改密码
		routerGroup.PUT("passwords", handler.ChangePassword)
		// 修改用户基础信息，包含昵称修改和 token 过期时间修改
		routerGroup.PUT("basics", handler.ChangeBasic)
		// 申请 API token
		routerGroup.GET("tokens", handler.ApplyAPIToken)
	}
}
