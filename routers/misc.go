package routers

import (
	"GoMD/controllers"
	"github.com/astaxie/beego"
)

func MiscRouter() {
	beego.Router("/login", &controllers.OtherController{}, "*:Login")   // 登陆
	beego.Router("/logout", &controllers.OtherController{}, "*:Logout") // 注销登陆
	//自定义页面
	beego.Router("/page/:title", &controllers.FrontendController{}, "get:Page")
}
