package controllers

import (
	"GoMD/models"
	"GoMD/tools"
	"github.com/astaxie/beego"
)

type FrontendController struct {
	beego.Controller
}

var theme = models.GetOneConfig("Theme")
var layout = theme+"/layout.html"

//首页逻辑
func (this *FrontendController) Index() {
	page, _ := this.GetInt64("page")
	pageSize := tools.StringToInt64(models.GetOneConfig("PageSize"))
	if page < 1 {
		page = 1
	}
	list,count := models.LimitArticleDisplay(page,pageSize)
	this.Data["paging"] = tools.CreatePaging(page, pageSize, count)
	this.Data["list"] = list
	this.Data["config"] = models.ConfigList()
	this.Data["page"] = page
	this.Layout = layout
	this.TplName = theme+"/index.html"
}

func (this *FrontendController) Page() {
	//文章查看页面
	id := this.GetString("id")
	if id == "" {
		id = "1"
	}
	article,_ := models.GetOneArticle(id)
	temp := *article
	label := models.SearchArticleCategory(temp[0].Cid)
	this.Data["article"] = article
	this.Data["label"] = label
	this.Data["config"] = models.ConfigList()
	this.Layout = layout
	this.TplName = theme+"/page.html"
}

func (this *FrontendController) Search() {
	//查询页面
	keywords := this.GetString("keywords")
	list,_ := models.Search(keywords)
	this.Data["list"] = list
	this.Data["config"] = models.ConfigList()
	this.Layout = layout
	this.TplName = theme+"/search.html"
}

func (this *FrontendController) About() {
	//关于 介绍页面
	this.Data["config"] = models.ConfigList()
	this.Layout = layout
	this.TplName = theme+"/about.html"
}

func (this *FrontendController) Archive() {
	//文章归档页面
	this.Data["config"] = models.ConfigList()
	data,_ := models.AllArticleList()
	this.Data["list"] = data
	this.Data["category"] = models.CategoryList()
	this.Layout = layout
	this.TplName = theme+"/archive.html"
}

func (this *FrontendController) Category() {
	// 输出分类下的所有文章
	id := this.GetString("id")
	this.Data["config"] = models.ConfigList()
	list,_ := models.GetCategoryArticle(id)
	this.Data["list"] = list
	this.Data["info"] = models.GetOneCategoryInfo(id)
	this.Layout = layout
	this.TplName = theme+"/category.html"
}