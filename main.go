package main

import (
	"GoMD/models"
	_ "GoMD/routers"
	"github.com/astaxie/beego"
)

func main() {
	//--------------------自定义 模板方法--------------------------
	beego.AddFuncMap("tags", Tags)                      //拆分标签
	beego.AddFuncMap("calc", Calc)                      //加减计算
	beego.AddFuncMap("markdown", MarkDown)              //将markdown输出为html
	beego.AddFuncMap("time", YMD)                       //将时间戳转成正常时间
	beego.AddFuncMap("sc", SiteConfig)                  //调取网站配置 直接抽调数据库配置字段
	beego.AddFuncMap("tableNum", TableNumber)           //获取表中的数据
	beego.AddFuncMap("notice", GetNotice)               //网站的公告
	beego.AddFuncMap("category", GetCategory)           //文章的分类信息
	beego.AddFuncMap("cn", GetAOfCategoryNumber)        //获取分类下的文章数量
	beego.AddFuncMap("comment", GetCommentNumber)       //获取分类下的评论数量
	beego.AddFuncMap("gavatar", GetGravatar)            //获取评论者gavatar头像
	beego.AddFuncMap("pna", PreOrNextAriticle)          //获取上一篇或者下一篇文章
	beego.AddFuncMap("newa", models.GetLimitNewArticle) //获取上最新文章列表
	beego.AddFuncMap("newc", models.GetLimitNewComment) //获取上最新文章评论
	beego.AddFuncMap("itu", IdToUuid)                   //根据id返回uuid
	beego.AddFuncMap("cl", models.CategoryList)         //返回分类列表
	beego.AddFuncMap("link", models.GetAllLink)         //返回链接列表
	beego.AddFuncMap("ed", EnumerateDate)               //返回单独的年月日

	//--------------------自定义页面
	beego.ErrorHandler("404", PageNotFound)

	//--------------------开放静态资源路径
	beego.SetStaticPath("/file", "file")
	beego.Run()
}
