package controllers

import (
	"GoMD/models"
	"GoMD/tools"
	"github.com/astaxie/beego"
	"strings"
)

/*
	comment 评论块
	sidebar 用户以及其他的信息块
*/

type FrontendController struct {
	beego.Controller
}

var theme = models.GetOneConfig("Theme")
var layout = theme + "/layout.html"

//首页逻辑
func (this *FrontendController) Index() {
	//检测伪静态
	config := models.ConfigList()
	var rule string
	if config["Rewrite"] == "1" {
		tmp := strings.Split(config["Repath"], "/")
		rule = tmp[1]
	} else {
		rule = "article"
	}

	page, _ := this.GetInt64("page")
	pageSize := tools.StringToInt64(models.GetOneConfig("PageSize"))
	if page < 1 {
		page = 1
	}
	list, count := models.LimitArticleDisplay(page, pageSize)
	this.Data["paging"] = tools.CreatePaging(page, pageSize, count)
	this.Data["list"] = list
	this.Data["config"] = config
	this.Data["page"] = page
	this.Data["rule"] = rule
	if config["Rewrite"] == "1" {
		this.Data["rewrite"] = "true"
	} else {
		this.Data["rewrite"] = "false"
	}
	this.Layout = layout
	this.TplName = theme + "/index.html"
	if theme == "QuietV1" {
		this.LayoutSections = make(map[string]string)
		this.LayoutSections["Sidebar"] = theme + "/sidebar.html"
	}
}

//文章查看页面
func (this *FrontendController) Article() {
	//检测伪静态
	config := models.ConfigList()
	var rule string
	name := this.GetString(":name")
	if config["Rewrite"] == "1" {
		rules := strings.Split(config["Repath"], "/")
		rule = rules[1]
		//规则不匹配 404
		if rule != this.GetString(":rule") {
			this.Redirect("/404", 302)
		}
		//拆分文件名
		tmp := strings.Split(name, ".")
		//获得名称
		name = tmp[0]
	} else {
		rule = "article"
	}
	article := models.GetOneArticle(name, "name")
	temp := *article
	label := models.SearchArticleCategory(temp[0].Cid)
	this.Data["id"] = temp[0].Id
	this.Data["article"] = article
	this.Data["rule"] = rule
	if config["Rewrite"] == "1" {
		this.Data["rewrite"] = "true"
	} else {
		this.Data["rewrite"] = "false"
	}
	this.Data["comments"] = models.GetArticleComments(temp[0].Id)
	this.Data["label"] = label
	this.Data["config"] = config
	this.Layout = layout
	this.TplName = theme + "/article.html"
	if theme == "QuietV1" {
		this.LayoutSections = make(map[string]string)
		this.LayoutSections["Sidebar"] = theme + "/sidebar.html"
		this.LayoutSections["Comment"] = theme + "/comment.html"
	}
}

func (this *FrontendController) Search() {
	//查询页面
	//检测伪静态
	config := models.ConfigList()
	var rule string
	if config["Rewrite"] == "1" {
		tmp := strings.Split(config["Repath"], "/")
		rule = tmp[1]
	} else {
		rule = "article"
	}
	this.Data["rule"] = rule
	if config["Rewrite"] == "1" {
		this.Data["rewrite"] = "true"
	} else {
		this.Data["rewrite"] = "false"
	}
	keywords := this.GetString("keywords")
	list, _ := models.Search(keywords)
	this.Data["keywords"] = keywords
	this.Data["list"] = list
	this.Data["config"] = models.ConfigList()
	this.Layout = layout
	this.TplName = theme + "/search.html"
}

//文章归档页面
func (this *FrontendController) Archive() {
	//检测伪静态
	config := models.ConfigList()
	var rule string
	if config["Rewrite"] == "1" {
		tmp := strings.Split(config["Repath"], "/")
		rule = tmp[1]
	} else {
		rule = "article"
	}
	this.Data["rule"] = rule
	if config["Rewrite"] == "1" {
		this.Data["rewrite"] = "true"
	} else {
		this.Data["rewrite"] = "false"
	}
	this.Data["config"] = config
	data, _ := models.AllArticleList()
	this.Data["list"] = data
	this.Data["category"] = models.CategoryList()
	this.Layout = layout
	this.TplName = theme + "/archive.html"
}

// 输出分类下的所有文章
func (this *FrontendController) Category() {
	//检测伪静态
	config := models.ConfigList()
	var rule string
	if config["Rewrite"] == "1" {
		tmp := strings.Split(config["Repath"], "/")
		rule = tmp[1]
	} else {
		rule = "article"
	}
	this.Data["data"] = models.CategoryInformation(this.GetString(":key"))
	this.Data["config"] = config
	this.Data["rule"] = rule
	if config["Rewrite"] == "1" {
		this.Data["rewrite"] = "true"
	} else {
		this.Data["rewrite"] = "false"
	}
	this.Layout = layout
	this.TplName = theme + "/category.html"
}
