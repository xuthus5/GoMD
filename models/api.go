package models

import (
	"GoMD/tools"
	"log"
)

/* ---------------------------

功能：用于统一的json数据发送与处理

------------------------------ */

//获取后台文章列表  后台文章页面调用该方法 返回一个json数据
func GetArticleJson() *[]DisplayArticle {
	list := []DisplayArticle{}
	err := dbx.Select(&list, "select article.id,article.title,article.author,category.name,article.tags,article.renew from article,category where article.cid=category.id and type=0 order by article.renew desc")
	if err != nil {
		log.Println(err.Error())
	}
	for i, v := range list {
		list[i].Renew = tools.UnixTimeToString(v.Renew)
	}
	return &list
}

//获取后台页面列表  后台文章页面调用该方法 返回一个json数据
func GetPageJson() *[]DisplayArticle {
	list := []DisplayArticle{}
	err := dbx.Select(&list, "select id,title,author,renew,name from article where type=1 order by article.renew desc")
	if err != nil {
		log.Println(err.Error())
	}
	for i, v := range list {
		list[i].Renew = tools.UnixTimeToString(v.Renew)
	}
	return &list
}

//获取后台分类列表
func GetCategoryJson() *[]Category {
	list := []Category{}
	err := dbx.Select(&list, "select * from category order by id desc")
	if err != nil {
		log.Println(err.Error())
	}
	return &list
}

//获取后台公告列表 后台notice页面调用该方法 返回一个json
func GetNoticeJson() *[]Notice {
	list := []Notice{}
	err := dbx.Select(&list, "select * from notice order by data desc")
	if err != nil {
		log.Println(err.Error())
	}
	for i, v := range list {
		list[i].Date = tools.UnixTimeToString(v.Date)
	}
	return &list
}

//获取后台文件列表
func GetFileJson() *[]Attachment {
	list := []Attachment{}
	err := dbx.Select(&list, "select * from attachment order by id desc")
	if err != nil {
		log.Println(err.Error())
	}
	for i, v := range list {
		list[i].Created = tools.UnixTimeToString(v.Created)
	}
	return &list
}
