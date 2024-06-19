package main

import (
	"gvb/gvb_server/core"
	"gvb/gvb_server/global"
	"gvb/gvb_server/routers"
)

func main() {
	core.InitConf()                //读取配置文件
	global.Log = core.InitLogger() //初始化日志
	global.DB = core.InitGorm()    //连接数据库
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("gvb_server运行在: %s", addr)
	router.Run(addr)
}
