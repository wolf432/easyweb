package routes

import (
	"easyweb/api/example"
	"github.com/gin-gonic/gin"
)

// SetExampleGroupRouter 定义示例路由
func SetExampleGroupRouter(router *gin.RouterGroup) {
	router.POST("/reg", example.Register)
}
