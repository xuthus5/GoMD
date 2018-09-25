package routers

import (
	"GoMD/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//前台界面
	beego.Router("/", &controllers.FrontendController{},"*:Index")
	beego.Router("/page", &controllers.FrontendController{},"*:Page")
	beego.Router("/search", &controllers.FrontendController{},"*:Search")
	beego.Router("/about", &controllers.FrontendController{},"*:About")
	beego.Router("/archive", &controllers.FrontendController{},"*:Archive")

	//后台界面
	beego.Router("/admin",&controllers.BackendController{},"get:Index") //后台主页
	beego.Router("/article",&controllers.BackendController{},"get:Article") //文章列表
	beego.Router("/article/add",&controllers.BackendController{},"get:ArticleAdd") //添加文章
	beego.Router("/article/update",&controllers.BackendController{},"get:ArticleUpdate") //更新文章
	beego.Router("/article/category",&controllers.BackendController{},"get:CategoryAdd") //添加标签
	beego.Router("/setting",&controllers.BackendController{},"get:Setting") //配置界面

	//api接口
	beego.Router("/api/article/list",&controllers.ApiController{},"get:ArticleList") //文章列表
	beego.Router("/api/article/add",&controllers.ApiController{},"post:ArticleAdd") //文章添加
	beego.Router("/api/article/update",&controllers.ApiController{},"post:ArticleUpdate") //文章更新
	beego.Router("/api/article/delete",&controllers.ApiController{},"get:ArticleDelete") //文章更新
	beego.Router("/api/category/add",&controllers.ApiController{},"post:CategoryAdd") //分类添加
	beego.Router("/api/site/config",&controllers.ApiController{},"post:SiteConfig") //网站配置


	//其他
	beego.Router("/login",&controllers.OtherController{},"*:Login") // 登陆
	beego.Router("/logout",&controllers.OtherController{},"*:Logout") // 注销登陆
}
