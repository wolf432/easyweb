package user

import "easyweb/pkg/validator"

type UserRegister struct {
	UserName string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Email    string `json:"email" form:"email" binding:"email"`
}

func (u UserRegister) GetMessages() validator.ValidatorMessageMap {
	return validator.ValidatorMessageMap{
		"username.required": "用户名不能为空",
		"password.required": "密码不能为空",
		"email.required":    "邮箱不能为空",
		"email.email":       "请填写正确的邮箱格式",
	}
}

type UserLogin struct {
	UserName string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (u UserLogin) GetMessages() validator.ValidatorMessageMap {
	return validator.ValidatorMessageMap{
		"userName.required": "用户名不能为空",
		"password.required": "密码不能为空",
	}
}
