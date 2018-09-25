package controllers

import (
	"GoMD/models"
	"github.com/astaxie/beego"
)

type BackendController struct {
	beego.Controller
}

/*
error  返回操作结果
0	没有错误
1	有错误
 */
type ResultData struct {
	Error 	int64
	Title 	string
	Msg 	string
}

//仪表盘 首页
func (this *BackendController) Index() {
	sess := this.GetSession("Username")
	if sess == nil{
		this.Redirect("/login",302)
	}
	this.Data["user"] = sess
	this.Layout = "admin/layout.html"
	this.TplName = "admin/index.html"
}

//文章管理列表
func (this *BackendController) Article() {
	sess := this.GetSession("Username")
	if sess == nil{
		this.Redirect("/login",302)
	}
	this.Data["user"] = sess
	this.Layout = "admin/layout.html"
	this.TplName = "admin/article.html"
}

//文章添加页面
func (this *BackendController) ArticleAdd() {
	sess := this.GetSession("Username")
	if sess == nil{
		this.Redirect("/login",302)
	}
	this.Data["user"] = sess
	this.Data["list"] = models.CategoryList()
	this.Layout = "admin/layout.html"
	this.TplName = "admin/article-add.html"
}
//文章更新页面
func (this *BackendController) ArticleUpdate() {
	sess := this.GetSession("Username")
	if sess == nil{
		this.Redirect("/login",302)
	}
	id := this.GetString("id")
	article,_ := models.GetOneArticle(id)
	this.Data["article"] = article
	this.Data["user"] = sess
	this.Data["category"] = models.CategoryList()
	this.Layout = "admin/layout.html"
	this.TplName = "admin/article-update.html"
}
//添加分类页面
func (this *BackendController) CategoryAdd()  {
	sess := this.GetSession("Username")
	if sess == nil{
		this.Redirect("/login",302)
	}
	this.Data["user"] = sess
	this.Data["list"] = models.CategoryList()
	this.Layout = "admin/layout.html"
	this.TplName = "admin/category.html"
}

func (this *BackendController) Setting() {
	//全局设置
	sess := this.GetSession("Username")
	if sess == nil{
		this.Redirect("/login",302)
	}
	this.Data["config"] = models.ConfigList()
	this.Data["user"] = sess
	this.Layout = "admin/layout.html"
	this.TplName = "admin/setting.html"
}