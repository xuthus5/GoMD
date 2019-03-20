package main

import (
	_ "GoMD/routers"
	"github.com/astaxie/beego"
)

func main() {
	//-----自定义模板方法
	customTemplateFunction()

	//-----自定义状态页面
	statusCode()

	//-----自定义静态资源路径
	staticSourcePath()

	//-----自定义日志配置
	logSetting()

	//开启热更新
	//beego.BConfig.Listen.Graceful=true

	//-----运行入口
	beego.Run()
}
