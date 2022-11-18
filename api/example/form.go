package example

import (
	"easyweb/pkg/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ExUser 验证字段说明文档:https://github.com/go-playground/validator
type ExUser struct {
	Name     string `form:"name" binding:"required" json:"name"`
	Password string `form:"password" binding:"required,min=2" json:"password"`
	Email    string `form:"email" binding:"required,email" json:"email"` //邮箱验证
	Phone    string `form:"phone" binding:"required,mobile"`
}

// GetMessages 定义表单验证的错误提示
func (u ExUser) GetMessages() validator.ValidatorMessageMap {
	//key 为结构体里form标签名+.+binding的验证标签
	return validator.ValidatorMessageMap{
		"name.required":     "用户名不能为空",
		"password.required": "密码不能为空",
		"password.min":      "密码长度不能小于2",
		"email.required":    "邮箱不能为空",
		"email.email":       "邮箱格式错误",
		"phone.mobile":      "手机格式错误",
	}
}

func Login(c *gin.Context) {
	c.String(200, "")
}

func Register(c *gin.Context) {
	var uF ExUser
	err := c.ShouldBind(&uF)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validator.GetErrorMsg(uF, err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
