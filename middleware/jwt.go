package middleware

import (
	"easyweb/global"
	pkgJwt "easyweb/pkg/jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func JWTAuth(iss string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Token")
		if tokenStr == "" {
			global.Log.Error("读取不到认证的token数据")
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "认证失败",
			})
			c.Abort()
			return
		}

		claims, err := pkgJwt.ParseToken(tokenStr, global.Cfg.Jwt.Secret)
		if err != nil {
			global.Log.Error("token解析失败", zap.Any("err", err))
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": -2,
				"msg":  "认证失败",
			})
			c.Abort()
			return
		}
		if claims.Issuer != iss {
			global.Log.Error("token认证失败,iss不相等")
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": -3,
				"msg":  "认证失败",
			})
			c.Abort()
			return
		}
		global.Log.Debug("token解析成功", zap.Any("claims", claims))
	}
}
