package main

import (
	"gvb/gvb_server/core"
	"gvb/gvb_server/flag"
	"gvb/gvb_server/global"
	"gvb/gvb_server/routers"
)

func main() {
	core.InitConf()                //读取配置文件
	global.Log = core.InitLogger() //初始化日志
	global.DB = core.InitGorm()    //连接数据库
	//命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("gvb_server运行在: %s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
