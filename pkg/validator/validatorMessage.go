// Package validator 设置验证方法返回自定义的错误信息
package validator

import (
	"easyweb/global"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type MyValidator interface {
	GetMessages() ValidatorMessageMap
}

// ValidatorMessageMap 保存错误字段和错误提示
// key为结构体里的字段名加上.加上标签名
type ValidatorMessageMap map[string]string

// GetErrorMsg 获取错误信息
func GetErrorMsg(request interface{}, err error) string {
	if _, isValidatorErrors := err.(validator.ValidationErrors); isValidatorErrors {
		_, isValidator := request.(MyValidator)

		for _, v := range err.(validator.ValidationErrors) {
			// 若 request 结构体实现 Validator 接口即可实现自定义错误信息
			if isValidator {
				if message, exist := request.(MyValidator).GetMessages()[v.Field()+"."+v.Tag()]; exist {
					return message
				}
			} else {
				global.Log.Debug("GetErrorMsg,key没找到", zap.Any("key", v.Field()+"."+v.Tag()))
			}
			return v.Error()
		}
	} else {
		return "err"
	}

	return "Parameter error"
}
