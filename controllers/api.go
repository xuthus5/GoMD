package controllers

import (
	"GoMD/models"
	"GoMD/tools"
	"github.com/astaxie/beego"
	"reflect"
	"strconv"
	"time"
)

type ApiController struct {
	beego.Controller
}

// 返回文章列表json数据数据格式
type AdminJsonData struct {
	Code  int                      `json:"code"`
	Count int                      `json:"count"`
	Msg   string                   `json:"msg"`
	Data  interface{} `json:"data"`
}

//返回后台文章列表 json数据类型返回 路由 /api/article/list
func (this *ApiController) ArticleList() {
	this.Data["json"] = &AdminJsonData{Code: 0, Count: 100, Msg: "", Data: models.GetArticleJson()}
	this.ServeJSON()
}
//返回后台分类列表 json数据类型返回 路由 /api/article/category/list
func (this *ApiController) CategoryList() {
	this.Data["json"] = &AdminJsonData{Code: 0, Count: 100, Msg: "", Data: models.GetCategoryJson()}
	this.ServeJSON()
}

//添加文章 数据校验  路由 /api/article/add
func (this *ApiController) ArticleAdd() {
	data := &models.Article{}
	info := &ResultData{}
	if err := this.ParseForm(data); err != nil {
		info = &ResultData{Error: 1, Title: "结果:", Msg: "接收表单数据出错！"}
	} else {
		data.Renew = tools.Int64ToString(time.Now().Unix())
		err := models.AddArticle(data)
		if err != nil {
			info = &ResultData{Error: 1, Title: "结果:", Msg: "数据库操作出错！"}
		} else {
			info = &ResultData{Error: 0, Title: "结果:", Msg: "发布成功！"}
		}
	}
	this.Data["json"] = info
	this.ServeJSON()
}

// 文章更新 数据校验  路由 /api/article/update
func (this *ApiController) ArticleUpdate() {
	id := this.GetString("id")
	data := &models.Article{}
	info := &ResultData{}
	if err := this.ParseForm(data); err != nil {
		info = &ResultData{Error: 1, Title: "结果:", Msg: "接收表单数据出错！"}
	} else {
		data.Id = tools.StringToInt(id)
		data.Renew = tools.Int64ToString(time.Now().Unix())
		err := models.UpdateArticle(data)
		if err != nil {
			info = &ResultData{Error: 1, Title: "结果:", Msg: "数据库操作出错！"}
		} else {
			info = &ResultData{Error: 0, Title: "结果:", Msg: "修改成功！"}
		}
	}
	this.Data["json"] = info
	this.ServeJSON()
}

// 文章删除  路由 /api/article/delete
func (this *ApiController) ArticleDelete() {
	info := &ResultData{}
	id, _ := strconv.Atoi(this.GetString("id"))
	err := models.DeleteArticle(id)
	if err != nil {
		info = &ResultData{Error: 1, Title: "结果:", Msg: "数据库操作出错！"}
	} else {
		info = &ResultData{Error: 0, Title: "结果:", Msg: "删除成功！"}
	}
	this.Data["json"] = info
	this.ServeJSON()
}

// 添加分类 路由 /api/article/category/add
func (this *ApiController) CategoryAdd() {
	name := this.GetString("name")
	key := this.GetString("key")
	description := this.GetString("description")
	info := &models.Taxonomy{Name: name, Key: key,Description:description}
	err := models.AddCategory(info)
	var data *ResultData
	if err != nil {
		data = &ResultData{Error: 1, Title: "结果:", Msg: "添加失败！"}
	} else {
		data = &ResultData{Error: 0, Title: "结果:", Msg: "添加成功！"}
	}
	this.Data["json"] = data
	this.ServeJSON()
}

// 分类删除 路由 /api/article/category/delete
func (this *ApiController) CategoryDelete() {
	info := &ResultData{}
	id, _ := strconv.Atoi(this.GetString("id"))
	err := models.DeleteCategory(id)
	if err != nil {
		info = &ResultData{Error: 1, Title: "结果:", Msg: "数据库操作出错！"}
	} else {
		info = &ResultData{Error: 0, Title: "结果:", Msg: "删除成功！"}
	}
	this.Data["json"] = info
	this.ServeJSON()
}

// 分类修改 路由 /api/article/category/update
func (this *ApiController) CategoryUpdate() {
	id := this.GetString("id")
	data := &models.Taxonomy{}
	info := &ResultData{}
	if err := this.ParseForm(data); err != nil {
		info = &ResultData{Error: 1, Title: "结果:", Msg: "接收表单数据出错！"}
	} else {
		data.Id = tools.StringToInt(id)
		err := models.UpdateCategory(data)
		if err != nil {
			info = &ResultData{Error: 1, Title: "结果:", Msg: "数据库操作出错！"}
		} else {
			info = &ResultData{Error: 0, Title: "结果:", Msg: "修改成功！"}
		}
	}
	this.Data["json"] = info
	this.ServeJSON()
}

// 网站设置页面  路由  /api/site/config
func (this *ApiController) SiteConfig() {
	submit := this.GetString("submit")
	info := &ResultData{}
	data := &models.ConfigOption{}
	if submit == "user" {
		config := &models.Config{Option: "Author", Value: this.GetString("username")}
		err := models.SiteConfig(config)
		config = &models.Config{Option: "Password", Value: this.GetString("password")}
		err1 := models.SiteConfig(config)
		if err != nil || err1 != nil {
			info = &ResultData{Error: 1, Title: "结果:", Msg: "出现未知错误！"}
		} else {
			info = &ResultData{Error: 0, Title: "结果:", Msg: "信息重置成功！"}
		}
	} else {
		if err := this.ParseForm(data); err != nil {
			info = &ResultData{Error: 1, Title: "结果:", Msg: "接收表单数据出错！"}
		} else {
			t := reflect.TypeOf(*data)  //类型
			v := reflect.ValueOf(*data) //值
			for i := 0; i < t.NumField(); i++ {
				config := &models.Config{Option: t.Field(i).Name, Value: v.Field(i).String()}
				err := models.SiteConfig(config)
				if err != nil {
					info = &ResultData{Error: 1, Title: "结果:", Msg: "出现未知错误！"}
				} else {
					info = &ResultData{Error: 0, Title: "结果:", Msg: "信息重置成功！"}
				}
			}
		}
	}
	this.Data["json"] = info
	this.ServeJSON()
}

// 公告数据接受 路由 /api/notice
func (this *ApiController) Notice() {
	notice := &models.Notice{Content:this.GetString("content"),Url:this.GetString("url"),Data:tools.Int64ToString(time.Now().Unix())}
	info := &ResultData{}
	if models.AddNotice(notice) != nil {
		info = &ResultData{Error: 1, Title: "结果:", Msg: "发布失败！"}
	}else{
		info = &ResultData{Error: 0, Title: "结果:", Msg: "发布成功！"}
	}
	this.Data["json"] = info
	this.ServeJSON()
}
// 公告列表 路由 /api/notice/list
func (this *ApiController) NoticeList() {
	this.Data["json"] = &AdminJsonData{Code: 0, Count: 100, Msg: "", Data: models.GetNoticeJson()}
	this.ServeJSON()
}