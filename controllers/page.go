package controllers

import (
	"GoMD/models"
	"strings"
)

/********************

新增单独的页面	诸如 关于我 友情连接

路由 ： /?:name
*********************/

func (this *FrontendController) Page() {
	//检测伪静态
	config := models.ConfigList()
	name := this.GetString(":name")
	if config["Rewrite"] == "1" {
		names := strings.Split(name, ".")
		name = names[0]
	}
	if config["Rewrite"] == "1" {
		this.Data["rewrite"] = "true"
	} else {
		this.Data["rewrite"] = "false"
	}
	article := models.GetOneArticle(name, "name")
	temp := *article
	this.Data["id"] = temp[0].Id
	this.Data["article"] = article
	this.Data["comments"] = models.GetArticleComments(temp[0].Id)
	this.Data["config"] = models.ConfigList()
	this.Layout = layout
	this.TplName = theme + "/page.html"
	if theme == "QuietV1" {
		this.LayoutSections = make(map[string]string)
		this.LayoutSections["Sidebar"] = theme + "/sidebar.html"
		this.LayoutSections["Comment"] = theme + "/comment.html"
	}
}
