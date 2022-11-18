package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

//定义错误码
const (
	// ErrDBDown 数据库挂掉了
	ErrDBDown = 1000

	// ErrDBNotFound 数据未找到
	ErrDBNotFound = 1001

	// ErrAuthFail JWT认证失败
	ErrAuthFail = 2000

	// ErrUser 用户模块错误
	ErrUser = 3000

	// ErrTag 标签模块错误
	ErrTag = 4000

	// ErrCategory 标签模块错误
	ErrCategory = 5000

	// ErrLink 链接模块错误
	ErrLink = 6000
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

// Ok 请求成功
func Ok(c *gin.Context) {
	Result(0, map[string]interface{}{}, "操作成功", c)
}

// OkWithMessage 请求成功，设置消息
func OkWithMessage(message string, c *gin.Context) {
	Result(0, map[string]interface{}{}, message, c)
}

// OkWithData 请求成功,设置返回数据
func OkWithData(data interface{}, c *gin.Context) {
	Result(0, data, "查询成功", c)
}

// OkWithDetailed 请求成功，设置消息和数据
func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(0, data, message, c)
}

// Fail 请求失败
func Fail(code int, c *gin.Context) {
	Result(code, map[string]interface{}{}, "操作失败", c)
}

// FailWithMessage 请求失败,设置错误码和错误信息
func FailWithMessage(code int, message string, c *gin.Context) {
	Result(code, map[string]interface{}{}, message, c)
}

// FailWithDetailed 请求失败，设置错误码、错误信息、数据
func FailWithDetailed(code int, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}
