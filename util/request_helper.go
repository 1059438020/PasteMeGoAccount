package util

import (
	"PasteMeGoAccount/vail"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

func ParseRequest(context *gin.Context, request interface{}) {
	err := context.ShouldBind(request)
	var errStr string
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errStr = vail.Translate(err.(validator.ValidationErrors))
		case *json.UnmarshalTypeError:
			unmarshalTypeError := err.(*json.UnmarshalTypeError)
			errStr = fmt.Errorf("%s 类型错误，期望类型 %s", unmarshalTypeError.Field, unmarshalTypeError.Type.String()).Error()
		default:
			errStr = errors.New("unknown error.").Error()
		}
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"error":   errStr,
			"message": "请求参数失败，请确认后重试",
		})
		log.Panic(errStr)
	}
}
