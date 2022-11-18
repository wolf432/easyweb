package tag

import "easyweb/pkg/validator"

// RequestTag 添加标签结构体
type RequestTag struct {
	Name string `json:"name" form:"name" binding:"required"`
}

// GetMessages 自定义的验证错误
func (t RequestTag) GetMessages() validator.ValidatorMessageMap {
	return validator.ValidatorMessageMap{
		"name.required": "标签名不能为空",
	}
}
