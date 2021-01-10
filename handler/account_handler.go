package handler

import (
	"PasteMeGoAccount/constant"
	"PasteMeGoAccount/model"
	"PasteMeGoAccount/request"
	"PasteMeGoAccount/util"
	"PasteMeGoAccount/view"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func RegisterAccount(context *gin.Context) {
	var (
		err error
		req request.Register
	)
	// 参数校验
	util.ParseRequest(context, &req)
	// 验证 Email 验证码
	//verification := HGet(constant.RedisEmailPrefix + req.Email, constant.RedisEmailVerification)
	//if verification != req.Verification {
	//	context.JSON(http.StatusOK, gin.H{
	//		"status":  http.StatusOK,
	//		"error":   "verification failed",
	//		"message": "邮箱验证码验证失败，请确认后重试",
	//	})
	//	return
	//}
	// 查询 Email 对应的 UID
	basic := model.BasicInfo{Email: req.Email}
	err = basic.Get()
	if err != nil {
		view.Error(context, err)
		return
	}
	// 该 Email 没有注册过需要初始化用户个人信息
	if basic.Uid == 0 {
		basic.Expiration = constant.DBTokenDefaultExpiration
		basic.Nickname = req.Username
		basic.State = constant.AccountActive
	}
	err = basic.Save()
	if err != nil {
		view.Error(context, err)
		return
	}
	// 初始化资源剩余访问次数
	// todo:bzy 初始化策源剩余访问次数
	// 校验账号是否已经存在
	uid := basic.Uid
	account := model.Account{Uid:uid}
	types := []uint8{constant.AccountUsername, constant.AccountEmail}
	accountExistFlag, err := account.ExistByUidAndType(types)
	if err != nil {
		view.Error(context, err)
		return
	}
	if accountExistFlag {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"error":   "account already exists",
			"message": "账号已被注册",
		})
		return
	}
	// 创建账号
	salt := strconv.FormatInt(time.Now().Unix(), 10)
	password := util.EncryptionPassword(req.Password, salt)
	accountUsername := model.Account{
		Uid:		  uid,
		Identifier:   req.Username,
		Credential:   password,
		Salt:         salt,
		IdentityType: constant.AccountUsername,
	}
	err = accountUsername.Save()
	if err != nil {
		view.Error(context, err)
		return
	}
	accountEmail := model.Account{
		Uid:		  uid,
		Identifier:   req.Email,
		Credential:   password,
		Salt:         salt,
		IdentityType: constant.AccountEmail,
	}
	err = accountEmail.Save()
	if err != nil {
		view.Error(context, err)
		return
	}
	view.Ok(context)
}

func FrozenAccount(context *gin.Context) {
	// 修改用户个人信息状态为禁用
}

func ChangePassword(context *gin.Context) {
	// 校验邮箱验证码并查询对应的 UID
	// 修改 UID 下所有非第三方的密码
}

func ChangeBasic(context *gin.Context) {
	// 修改 UID 对应的信息
}

func ApplyAPIToken(context *gin.Context) {
	// 校验邮箱验证码
	// 验证返回当前 Token
}

func SendEmail(context *gin.Context) {

}
