package vail

import (
	"github.com/gin-gonic/gin/binding"
	zhongwen "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

var trans ut.Translator

func InitValidator() {
	// 中文翻译
	zh := zhongwen.New()
	uni := ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		// 验证器注册翻译器
		zh_translations.RegisterDefaultTranslations(v, trans)
		// 自定义验证方法
		v.RegisterValidation("checkEmail", checkEmail)
	}
}

func Translate(errs validator.ValidationErrors) string {
	var errList []string
	for _, e := range errs {
		// can translate each error one at a time.
		errList = append(errList, e.Translate(trans))
	}
	return strings.Join(errList, "|")
}

func checkEmail(fl validator.FieldLevel) bool {
	//email := fl.Field().String()
	//re := `^1[3456789]\d{9}$`
	//r := regexp.MustCompile(re)
	//return r.MatchString(email)
	return true
}