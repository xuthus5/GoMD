package controllers

import (
	"GoMD/models"
	"github.com/astaxie/beego"
)

type BackendController struct {
	beego.Controller
}

//*******************控制台*************************

//仪表盘 首页~概要		路由：/admin
func (this *BackendController) Index() {
	master := this.GetSession("master")
	if master == nil {
		this.Redirect("/login", 302)
	}
	this.Data["master"] = master
	this.Layout = "admin/layout.html"
	this.TplName = "admin/index.html"
}

//外观设置			路由： /admin/style
func (this *BackendController) Style() {
	master := this.GetSession("master")
	if master == nil {
		this.Redirect("/login", 302)
	}
	this.Data["master"] = master
	this.Layout = "admin/layout.html"
	this.TplName = "admin/index.html"
}

//网站数据备份			路由： /admin/backup
func (this *BackendController) Backup() {
	master := this.GetSession("master")
	if master == nil {
		this.Redirect("/login", 302)
	}
	this.Data["master"] = master
	this.Layout = "admin/layout.html"
	this.TplName = "admin/backup.html"
}

//*******************撰写*************************

//文章添加页面				路由： /admin/article/add
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

//页面添加				路由： /admin/page/add
func (this *BackendController) PageAdd() {
	master := this.GetSession("master")
	if master == nil {
		this.Redirect("/login", 302)
	}
	this.Data["master"] = master
	this.Layout = "admin/layout.html"
	this.TplName = "admin/page/add.html"
}

//*******************管理*************************
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

//页面管理列表
func (this *BackendController) Page() {
	master := this.GetSession("master")
	if master == nil {
		this.Redirect("/login", 302)
	}
	this.Data["master"] = master
	this.Layout = "admin/layout.html"
	this.TplName = "admin/page/list.html"
}

// 公告管理
func (this *BackendController) Notice() {
	master := this.GetSession("master")
	if master == nil {
		this.Redirect("/login", 302)
	}
	this.Data["master"] = master
	this.Layout = "admin/layout.html"
	this.TplName = "admin/manager/notice.html"
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
	this.TplName = "admin/manager/attachment.html"
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
	this.TplName = "admin/manager/link.html"
}

//菜单页面
func (this *BackendController) Menu() {
	master := this.GetSession("master")
	if master == nil {
		this.Redirect("/login", 302)
	}
	this.Data["master"] = master
	this.Data["config"] = models.ConfigList()
	this.Layout = "admin/layout.html"
	this.TplName = "admin/manager/menu.html"
}

//评论页面
func (this *BackendController) Comment() {
	master := this.GetSession("master")
	if master == nil {
		this.Redirect("/login", 302)
	}
	this.Data["master"] = master
	this.Data["config"] = models.ConfigList()
	this.Layout = "admin/layout.html"
	this.TplName = "admin/manager/comment.html"
}

//*******************设置*************************
//全局设置
func (this *BackendController) Setting() {
	master := this.GetSession("master")
	if master == nil {
		this.Redirect("/login", 302)
	}
	this.Data["master"] = master
	this.Data["config"] = models.ConfigList()
	this.Layout = "admin/layout.html"
	this.TplName = "admin/setting/website.html"
}

// 固定链接设置
func (this *BackendController) PathRewrite()  {
	master := this.GetSession("master")
	if master == nil {
		this.Redirect("/login", 302)
	}
	this.Data["master"] = master
	this.Data["config"] = models.ConfigList()
	this.Layout = "admin/layout.html"
	this.TplName = "admin/setting/path.html"
}
//*******************其他必要的页面*************************

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

//页面更新页面
func (this *BackendController) PageUpdate() {
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
	this.TplName = "admin/page/update.html"
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

//修改链接页面
func (this *BackendController) LinkUpdate() {
	master := this.GetSession("master")
	if master == nil {
		this.Redirect("/login", 302)
	}
	this.Data["master"] = master
	this.Data["link"] = models.GetOneLinkInfo(this.GetString("id"))
	this.Layout = "admin/layout.html"
	this.TplName = "admin/manager/link-update.html"
}

//修改菜单节点页面
func (this *BackendController) MenuNodeUpdate() {
	master := this.GetSession("master")
	if master == nil {
		this.Redirect("/login", 302)
	}
	this.Data["master"] = master
	this.Data["link"] = models.GetOneLinkInfo(this.GetString("id"))
	this.Layout = "admin/layout.html"
	this.TplName = "admin/manager/menu-update.html"
}
