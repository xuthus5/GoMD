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