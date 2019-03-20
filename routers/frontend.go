package routers

import (
	"GoMD/controllers"
	"github.com/astaxie/beego"
)

func FrontendRouter() {
	beego.Router("/", &controllers.FrontendController{}, "*:Index")
	beego.Router("/:name", &controllers.FrontendController{}, "*:Page")
	beego.Router("/article/:name", &controllers.FrontendController{}, "*:Article")
	beego.Router("/search", &controllers.FrontendController{}, "*:Search")
	beego.Router("/archive", &controllers.FrontendController{}, "*:Archive")
	beego.Router("/category/:key", &controllers.FrontendController{}, "*:Category")
}
