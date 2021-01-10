package router

import (
	"PasteMeGoAccount/handler"
	"github.com/gin-gonic/gin"
)

func LoadOtherRouter(ginEngine *gin.Engine) {
	routerGroup := ginEngine.Group("other")
	{
		// 向指定邮箱发送消息
		routerGroup.POST("email", handler.RegisterAccount)
	}
}
