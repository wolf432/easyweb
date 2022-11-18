package user

import (
	"easyweb/global"
	"easyweb/pkg/response"
	"easyweb/pkg/validator"
	"easyweb/service/user"
	userReq "easyweb/service/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func Login(c *gin.Context) {
	var userLogin userReq.UserLogin
	if err := c.ShouldBind(&userLogin); err != nil {
		global.Log.Debug("用户登录接收参数错误:", zap.Error(err))
		response.FailWithMessage(response.ErrAuthFail, validator.GetErrorMsg(userLogin, err), c)
		return
	}
	userInfo, err := user.Login(c, &userLogin)
	if err != nil {
		global.Log.Debug("用户登录失败:", zap.Error(err))
		response.FailWithMessage(response.ErrUser, err.Error(), c)
		return
	}
	response.OkWithData(userInfo, c)
}

func Register(c *gin.Context) {
	var newUser userReq.UserRegister
	if err := c.ShouldBind(&newUser); err != nil {
		global.Log.Debug("用户注册接收参数错误:", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  validator.GetErrorMsg(newUser, err),
		})
		return
	}
	userInfo, err := user.Register(c, &newUser)
	if err != nil {
		global.Log.Debug("用户注册失败:", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	global.Log.Debug("用户注册成功:", zap.Any("info", userInfo))
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "注册成功",
		"data": userInfo,
	})
}
