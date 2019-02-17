package routers

import (
	"GoMD/controllers"
	"github.com/astaxie/beego"
)

func BackendRouter() {
	beego.Router("/admin", &controllers.BackendController{}, "get:Index")                                  //后台主页
	beego.Router("/admin/notice", &controllers.BackendController{}, "get:Notice")                          //公告功能
	beego.Router("/admin/article", &controllers.BackendController{}, "get:Article")                        //文章列表
	beego.Router("/admin/article/add", &controllers.BackendController{}, "get:ArticleAdd")                 //添加文章
	beego.Router("/admin/article/update", &controllers.BackendController{}, "get:ArticleUpdate")           //更新文章
	beego.Router("/admin/article/category", &controllers.BackendController{}, "get:CategoryAdd")           //添加标签
	beego.Router("/admin/article/category/update", &controllers.BackendController{}, "get:CategoryUpdate") //修改标签
	beego.Router("/admin/setting", &controllers.BackendController{}, "get:Setting")                        //配置界面
	beego.Router("/admin/file", &controllers.BackendController{}, "get:Attachment")                        //附件管理界面
	beego.Router("/admin/link", &controllers.BackendController{}, "get:Link")                              //链接管理界面
	beego.Router("/admin/link/update", &controllers.BackendController{}, "get:LinkUpdate")                 //链接修改界面
	beego.Router("/admin/page/add", &controllers.BackendController{}, "get:PageAdd")                       //添加页面
	beego.Router("/admin/page/update", &controllers.BackendController{}, "get:PageUpdate")                 //修改页面
	beego.Router("/admin/page", &controllers.BackendController{}, "get:Page")                              //页面列表管理
	beego.Router("/admin/menu", &controllers.BackendController{}, "get:Menu")                              //菜单管理
	beego.Router("/admin/menu/update", &controllers.BackendController{}, "get:MenuNodeUpdate")             //菜单节点修改
	beego.Router("/admin/comment", &controllers.BackendController{}, "get:Comment")                        //评论页面
	beego.Router("/admin/style", &controllers.BackendController{}, "get:Style")                            //主题样式配置页面
	beego.Router("/admin/backup", &controllers.BackendController{}, "get:Backup")                          //备份页面
}
