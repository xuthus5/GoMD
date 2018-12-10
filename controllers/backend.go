package controllers

import (
	"GoMD/models"
	"github.com/astaxie/beego"
)

type BackendController struct {
	beego.Controller
}

//仪表盘 首页
func (this *BackendController) Index() {
	master := this.GetSession("master")
	if master == nil {
		this.Redirect("/login", 302)
	}
	this.Data["master"] = master
	this.Layout = "admin/layout.html"
	this.TplName = "admin/index.html"
}

// 公告管理
func (this *BackendController) Notice() {
	master := this.GetSession("master")
	if master == nil {
		this.Redirect("/login", 302)
	}
	this.Data["master"] = master
	this.Layout = "admin/layout.html"
	this.TplName = "admin/notice.html"
}

//文章管理列表
func (this *BackendController) Article() {
	master := this.GetSession("master")
	if master == nil {
		this.Redirect("/login", 302)
	}
	this.Data["master"] = master
	this.Layout = "admin/layout.html"
	this.TplName = "admin/article/list.html"
}

//文章添加页面
func (this *BackendController) ArticleAdd() {
	master := this.GetSession("master")
	if master == nil {
		this.Redirect("/login", 302)
	}
	this.Data["master"] = master
	this.Data["list"] = models.CategoryList()
	this.Layout = "admin/layout.html"
	this.TplName = "admin/article/add.html"
}

//文章更新页面
func (this *BackendController) ArticleUpdate() {
	master := this.GetSession("master")
	if master == nil {
		this.Redirect("/login", 302)
	}
	this.Data["master"] = master
	id := this.GetString("id")
	article := models.GetOneArticle(id, "id")
	this.Data["article"] = article
	this.Data["category"] = models.CategoryList()
	this.Layout = "admin/layout.html"
	this.TplName = "admin/article/update.html"
}

//添加分类页面
func (this *BackendController) CategoryAdd() {
	master := this.GetSession("master")
	if master == nil {
		this.Redirect("/login", 302)
	}
	this.Data["master"] = master
	this.Data["list"] = models.CategoryList()
	this.Layout = "admin/layout.html"
	this.TplName = "admin/article/category.html"
}

//修改分类页面
func (this *BackendController) CategoryUpdate() {
	master := this.GetSession("master")
	if master == nil {
		this.Redirect("/login", 302)
	}
	this.Data["master"] = master
	this.Data["category"] = models.GetOneCategoryInfo(this.GetString("id"), "id")
	this.Layout = "admin/layout.html"
	this.TplName = "admin/article/category-update.html"
}

//全局设置
func (this *BackendController) Setting() {
	master := this.GetSession("master")
	if master == nil {
		this.Redirect("/login", 302)
	}
	this.Data["master"] = master
	this.Data["config"] = models.ConfigList()
	this.Layout = "admin/layout.html"
	this.TplName = "admin/setting.html"
}

// 附件管理
func (this *BackendController) Attachment() {
	master := this.GetSession("master")
	if master == nil {
		this.Redirect("/login", 302)
	}
	this.Data["master"] = master
	this.Data["config"] = models.ConfigList()
	this.Layout = "admin/layout.html"
	this.TplName = "admin/attachment.html"
}

// 链接管理
func (this *BackendController) Link() {
	master := this.GetSession("master")
	if master == nil {
		this.Redirect("/login", 302)
	}
	this.Data["master"] = master
	this.Data["config"] = models.ConfigList()
	this.Layout = "admin/layout.html"
	this.TplName = "admin/link.html"
}

//修改链接页面
func (this *BackendController) LinkUpdate() {
	master := this.GetSession("master")
	if master == nil {
		this.Redirect("/login", 302)
	}
	this.Data["master"] = master
	this.Data["link"] = models.GetOneLinkInfo(this.GetString("id"))
	this.Layout = "admin/layout.html"
	this.TplName = "admin/link-update.html"
}
