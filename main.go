package main

import (
	"easyweb/global"
	"easyweb/initialize"
)

func main() {

	initialize.InitViper()
	global.Log = initialize.InitLog()
	global.DB = initialize.InitDb()
	global.REdis = initialize.InitRedis()
	initialize.InitializeValidator()

	defer func() {
		//程序关闭前，是否数据库链接
		if global.DB != nil {
			db, _ := global.DB.DB()
			db.Close()
		}
	}()

	//启动Gin
	initialize.RunServer()
}
