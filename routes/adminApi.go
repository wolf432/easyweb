package routes

import (
	"easyweb/api/v1/category"
	"easyweb/api/v1/link"
	"easyweb/api/v1/tag"
	"github.com/gin-gonic/gin"
)

// SetApiAuthGroupRouter   定义需要验证的分组路由
func SetApiAuthGroupRouter(router *gin.RouterGroup) {
	//========== 标签路由
	router.POST("/tag", tag.Create)
	router.DELETE("/tag/:id", tag.Delete)
	router.PUT("/tag/:id", tag.Update)

	//========== 分类路由
	router.POST("/category", category.Create)
	router.DELETE("/category/:id", category.Delete)
	router.PUT("/category", category.Update)

	//========== 链接路由
	router.POST("/link", link.Create)
	router.DELETE("/link/:id", link.Delete)
	router.PUT("/link", link.Update)
}
