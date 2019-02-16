package routers

import (
	"GoMD/controllers"
	"github.com/astaxie/beego"
)

func ApiRouter() {
	beego.Router("/api/article/list", &controllers.ApiController{}, "get:ArticleList")           //文章列表
	beego.Router("/api/article/category/list", &controllers.ApiController{}, "get:CategoryList") //文章列表
	beego.Router("/api/article/add", &controllers.ApiController{}, "post:ArticleAdd")            //文章添加
	beego.Router("/api/article/update", &controllers.ApiController{}, "post:ArticleUpdate")      //文章更新
	beego.Router("/api/article/delete", &controllers.ApiController{}, "get:ArticleDelete")       //文章更新

	beego.Router("/api/article/category/add", &controllers.ApiController{}, "post:CategoryAdd")       //分类添加
	beego.Router("/api/article/category/delete", &controllers.ApiController{}, "get:CategoryDelete")  //分类删除
	beego.Router("/api/article/category/update", &controllers.ApiController{}, "post:CategoryUpdate") //分类修改

	beego.Router("/api/page/list", &controllers.ApiController{}, "get:PageList")      //页面列表
	beego.Router("/api/page/add", &controllers.ApiController{}, "post:PageAdd")       //页面添加
	beego.Router("/api/page/update", &controllers.ApiController{}, "post:PageUpdate") //页面更新
	beego.Router("/api/page/delete", &controllers.ApiController{}, "get:PageDelete")  //页面删除

	beego.Router("/api/site/config", &controllers.ApiController{}, "post:SiteConfig") //网站配置

	beego.Router("/api/notice", &controllers.ApiController{}, "post:Notice")         //网站公告
	beego.Router("/api/notice/list", &controllers.ApiController{}, "get:NoticeList") //网站公告

	beego.Router("/api/file/upload", &controllers.ApiController{}, "post:FileUpload") //文件上传
	beego.Router("/api/file/list", &controllers.ApiController{}, "get:FileList")      //文件上传
	beego.Router("/api/file/delete", &controllers.ApiController{}, "get:FileDelete")  //文件删除

	beego.Router("/api/comment/add", &controllers.ApiController{}, "post:CommentAdd") //评论添加

	beego.Router("/api/link/delete", &controllers.ApiController{}, "get:LinkDelete")  //链接删除
	beego.Router("/api/link/add", &controllers.ApiController{}, "post:LinkAdd")       //链接添加
	beego.Router("/api/link/list", &controllers.ApiController{}, "get:LinkList")      //链接列表
	beego.Router("/api/link/update", &controllers.ApiController{}, "post:LinkUpdate") //链接修改

	beego.Router("/api/menu/delete", &controllers.ApiController{}, "get:MenuNodeDelete")  //菜单节点删除
	beego.Router("/api/menu/add", &controllers.ApiController{}, "post:MenuNodeAdd")       //菜单节点添加
	beego.Router("/api/menu/list", &controllers.ApiController{}, "get:MenuList")          //菜单节点列表
	beego.Router("/api/menu/update", &controllers.ApiController{}, "post:MenuNodeUpdate") //菜单节点修改

	beego.Router("/api/comment/review", &controllers.ApiController{}, "get:ReviewComment") //待审核评论列表
	beego.Router("/api/comment/adopt", &controllers.ApiController{}, "get:AdoptComment")   //通过审核的评论列表
	beego.Router("/api/comment/delete", &controllers.ApiController{}, "get:DeleteComment") //删除评论
	beego.Router("/api/comment/update", &controllers.ApiController{}, "get:UpdateComment") //通过评论操作
}
