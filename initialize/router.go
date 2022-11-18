package initialize

import (
	"context"
	"easyweb/global"
	"easyweb/middleware"
	"easyweb/routes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())

	r.GET("/ping", func(c *gin.Context) {
		time.Sleep(time.Second * 10)
		c.String(http.StatusOK, "ping")
	})

	//示例路由
	exampleGroup := r.Group("/example")
	routes.SetExampleGroupRouter(exampleGroup)

	// 注册 api 分组路由,不需要验证
	v1Group := r.Group("/v1")
	routes.SetApiGroupRouter(v1Group)

	v1AuthGroup := r.Group("/v1")
	v1AuthGroup.Use(middleware.JWTAuth("web"))
	routes.SetApiAuthGroupRouter(v1AuthGroup)

	return r
}

// RunServer 启动服务器
func RunServer() {
	r := setupRouter()

	srv := &http.Server{
		Addr:    ":" + global.Cfg.App.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Log.Fatal("listen:", zap.Any("error", err))
		}
	}()

	// 等待中断信号优雅的关闭服务器
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.Log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.Log.Fatal("Server Shutdown:", zap.Any("error", err))
	}
	global.Log.Info("Server exiting")
}
