package routers

import (
	"GoMD/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//前台界面
	beego.Router("/", &controllers.FrontendController{}, "*:Index")
	beego.Router("/article/:uuid", &controllers.FrontendController{}, "*:Article")
	beego.Router("/search", &controllers.FrontendController{}, "*:Search")
	beego.Router("/about", &controllers.FrontendController{}, "*:About")
	beego.Router("/archive", &controllers.FrontendController{}, "*:Archive")
	beego.Router("/category/:key", &controllers.FrontendController{}, "*:Category")

	//后台界面
	beego.Router("/admin", &controllers.BackendController{}, "get:Index")                                  //后台主页
	beego.Router("/admin/notice", &controllers.BackendController{}, "get:Notice")                          //公告功能
	beego.Router("/admin/article", &controllers.BackendController{}, "get:Article")                        //文章列表
	beego.Router("/admin/article/add", &controllers.BackendController{}, "get:ArticleAdd")                 //添加文章
	beego.Router("/admin/article/update", &controllers.BackendController{}, "get:ArticleUpdate")           //更新文章
	beego.Router("/admin/article/category", &controllers.BackendController{}, "get:CategoryAdd")           //添加标签
	beego.Router("/admin/article/category/update", &controllers.BackendController{}, "get:CategoryUpdate") //修改标签
	beego.Router("/admin/setting", &controllers.BackendController{}, "get:Setting")                        //配置界面
	beego.Router("/admin/attachment", &controllers.BackendController{}, "get:Attachment")                  //附件管理界面
	beego.Router("/admin/link", &controllers.BackendController{}, "get:Link")                              //链接管理界面
	beego.Router("/admin/link/update", &controllers.BackendController{}, "get:LinkUpdate")                 //链接修改界面

	//api接口
	beego.Router("/api/article/list", &controllers.ApiController{}, "get:ArticleList")                //文章列表
	beego.Router("/api/article/category/list", &controllers.ApiController{}, "get:CategoryList")      //文章列表
	beego.Router("/api/article/add", &controllers.ApiController{}, "post:ArticleAdd")                 //文章添加
	beego.Router("/api/article/update", &controllers.ApiController{}, "post:ArticleUpdate")           //文章更新
	beego.Router("/api/article/delete", &controllers.ApiController{}, "get:ArticleDelete")            //文章更新
	beego.Router("/api/article/category/add", &controllers.ApiController{}, "post:CategoryAdd")       //分类添加
	beego.Router("/api/article/category/delete", &controllers.ApiController{}, "get:CategoryDelete")  //分类删除
	beego.Router("/api/article/category/update", &controllers.ApiController{}, "post:CategoryUpdate") //分类修改
	beego.Router("/api/site/config", &controllers.ApiController{}, "post:SiteConfig")                 //网站配置
	beego.Router("/api/notice", &controllers.ApiController{}, "post:Notice")                          //网站配置
	beego.Router("/api/notice/list", &controllers.ApiController{}, "get:NoticeList")                  //网站配置
	beego.Router("/api/file/upload", &controllers.ApiController{}, "post:FileUpload")                 //文件上传
	beego.Router("/api/file/list", &controllers.ApiController{}, "get:FileList")                      //文件上传
	beego.Router("/api/file/delete", &controllers.ApiController{}, "get:FileDelete")                  //文件删除
	beego.Router("/api/comment/add", &controllers.ApiController{}, "post:CommentAdd")                 //评论添加
	beego.Router("/api/link/delete", &controllers.ApiController{}, "get:LinkDelete")                  //链接删除
	beego.Router("/api/link/add", &controllers.ApiController{}, "post:LinkAdd")                       //链接添加
	beego.Router("/api/link/list", &controllers.ApiController{}, "get:LinkList")                      //链接列表
	beego.Router("/api/link/update", &controllers.ApiController{}, "post:LinkUpdate")                 //链接修改

	//其他
	beego.Router("/login", &controllers.OtherController{}, "*:Login")   // 登陆
	beego.Router("/logout", &controllers.OtherController{}, "*:Logout") // 注销登陆
	//自定义页面
	beego.Router("/page/:title",&controllers.FrontendController{},"get:Page")
}
