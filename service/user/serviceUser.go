package user

import (
	"easyweb/global"
	"easyweb/model"
	"easyweb/pkg/hash"
	"easyweb/pkg/jwt"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os/user"
)

//Login 用户登录操作
func Login(c *gin.Context, userInfo *UserLogin) (*ResponseUserInfo, error) {
	if userInfo == nil {
		return nil, errors.New("请填写必要信息后再登录")
	}
	var user model.User
	err := global.DB.Where("username = ?", userInfo.UserName).Take(&user).Error
	if err == nil {
		//密码错误
		if false == hash.CheckPasswordHash(userInfo.Password, user.Password) {
			return nil, errors.New("用户名或密码错误")
		}
	}
	//找不到数据
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户名或密码错误")
	}
	//生成jwt的token
	token, _ := jwt.CreateToken("web", user, global.Cfg.Jwt.Secret, global.Cfg.Jwt.JwtTtl)
	return &ResponseUserInfo{
		Uid:      user.Id,
		UserName: user.UserName,
		Email:    user.Email,
		Token:    token,
	}, nil
}

//Register 用户注册
func Register(c *gin.Context, userInfo *UserRegister) (*ResponseUserInfo, error) {
	var user model.User
	err := global.DB.Where("username = ?", userInfo.UserName).Take(&user).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户已注册，不可以重复注册")
	}
	user.UserName = userInfo.UserName
	user.Email = userInfo.Email
	user.Password, err = hash.HashPassword(userInfo.Password)
	if err != nil {
		global.Log.Error("用户注册加密密码失败:", zap.Error(err))
		return nil, errors.New("注册失败，请重新尝试")
	}
	if err := global.DB.Create(&user).Error; err != nil {
		global.Log.Error("用户注册失败", zap.Error(err))
		return nil, errors.New("注册失败，请重新尝试")
	}
	//生成jwt的token
	token, _ := jwt.CreateToken("web", user, global.Cfg.Jwt.Secret, global.Cfg.Jwt.JwtTtl)
	return &ResponseUserInfo{
		Uid:      user.Id,
		UserName: user.UserName,
		Email:    user.Email,
		Token:    token,
	}, nil
}

// Delete 删除用户
func Delete(c *gin.Context, uid uint) error {
	var user user.User
	err := global.DB.First(&user, uid).Error
	if err != nil {
		return errors.New("没找到该用户")
	}
	return global.DB.Where("id = ? ", uid).Delete(&user).Error
}
