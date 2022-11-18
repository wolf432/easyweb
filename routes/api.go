package routes

import (
	"easyweb/api/v1/category"
	"easyweb/api/v1/link"
	"easyweb/api/v1/tag"
	"easyweb/api/v1/user"
	"github.com/gin-gonic/gin"
)

// SetApiGroupRouter   定义不需要验证的分组路由
func SetApiGroupRouter(router *gin.RouterGroup) {
	//========== 标签路由
	router.GET("tag/:id", tag.GetInfo)
	router.GET("tags", tag.GetList)

	//========== 分类路由
	router.GET("category/:id", category.GetInfo)
	router.GET("categories", category.GetList)

	//========== 链接路由
	router.GET("link/:id", link.GetInfo)
	router.GET("links", link.GetList)

	//========== 用户路由
	router.POST("user/register", user.Register)
	router.POST("user/login", user.Login)
}
