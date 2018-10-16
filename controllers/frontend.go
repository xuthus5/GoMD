package controllers

import (
	"GoMD/models"
	"GoMD/tools"
	"github.com/astaxie/beego"
	"strconv"
)

type FrontendController struct {
	beego.Controller
}

var theme = models.GetOneConfig("Theme")
var layout = theme + "/layout.html"

//首页逻辑
func (this *FrontendController) Index() {
	page, _ := this.GetInt64("page")
	pageSize := tools.StringToInt64(models.GetOneConfig("PageSize"))
	if page < 1 {
		page = 1
	}
	list, count := models.LimitArticleDisplay(page, pageSize)
	this.Data["paging"] = tools.CreatePaging(page, pageSize, count)
	this.Data["list"] = list
	this.Data["config"] = models.ConfigList()
	this.Data["page"] = page
	this.Layout = layout
	this.TplName = theme + "/index.html"
}

func (this *FrontendController) Article() {
	//文章查看页面
	id := this.GetString(":uuid")
	article := models.GetOneArticle(id,"uuid")
	temp := *article
	label := models.SearchArticleCategory(temp[0].Cid)
	this.Data["article"] = article
	this.Data["label"] = label
	this.Data["config"] = models.ConfigList()
	this.Layout = layout
	this.TplName = theme + "/page.html"
}

func (this *FrontendController) Search() {
	//查询页面
	keywords := this.GetString("keywords")
	list, _ := models.Search(keywords)
	this.Data["list"] = list
	this.Data["config"] = models.ConfigList()
	this.Layout = layout
	this.TplName = theme + "/search.html"
}

func (this *FrontendController) About() {
	//关于 介绍页面
	this.Data["config"] = models.ConfigList()
	this.Layout = layout
	this.TplName = theme + "/about.html"
}

func (this *FrontendController) Archive() {
	//文章归档页面
	this.Data["config"] = models.ConfigList()
	data, _ := models.AllArticleList()
	this.Data["list"] = data
	this.Data["category"] = models.CategoryList()
	this.Layout = layout
	this.TplName = theme + "/archive.html"
}

func (this *FrontendController) Category() {
	// 输出分类下的所有文章
	key := this.GetString(":key")
	this.Data["config"] = models.ConfigList()
	info := models.GetOneCategoryInfo(key,"key")
	temp := *info
	list, _ := models.GetCategoryArticle(strconv.Itoa(temp[0].Id))
	this.Data["list"] = list
	// 输出该分类的信息
	this.Data["info"] = info
	this.Layout = layout
	this.TplName = theme + "/category.html"
}
