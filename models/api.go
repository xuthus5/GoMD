package models

import "GoMD/tools"

/* ---------------------------

功能：用于统一的json数据发送与处理

------------------------------ */


//获取后台文章列表  后台文章页面调用该方法 返回一个json数据
func GetArticleJson() *[]DisplayArticle{
	list := []DisplayArticle{}
	err := dbx.Select(&list, "select article.id,article.title,article.author,taxonomy.name,article.tags,article.renew from article,taxonomy where article.cid=taxonomy.id order by article.renew desc")
	if err != nil {
		panic(err.Error())
	}
	for i,v:= range list{
		list[i].Renew = tools.UnixTimeToString(v.Renew)
	}
	return &list
}

//获取后台分类列表
func GetCategoryJson() *[]Taxonomy{
	list := []Taxonomy{}
	err := dbx.Select(&list, "select * from taxonomy order by id desc")
	if err != nil {
		panic(err.Error())
	}
	return &list
}

//获取后台公告列表 后台notice页面调用该方法 返回一个json
func GetNoticeJson() *[]Notice {
	list := []Notice{}
	err := dbx.Select(&list,"select * from notice order by data desc")
	if err != nil {
		panic(err.Error())
	}
	for i,v:= range list{
		list[i].Data = tools.UnixTimeToString(v.Data)
	}
	return &list
}