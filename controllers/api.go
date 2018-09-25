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
type ArticleData struct {
	Code  int           `json:"code"`
	Count int           `json:"count"`
	Msg   string        `json:"msg"`
	Data  *[]models.DisplayArticle `json:"data"`
}
//返回后台文章列表 json数据类型返回
func (this *ApiController) ArticleList() {
	list := models.GetArticleJson()
	data := ArticleData{Code: 0, Count: 100, Msg: "", Data: list}
	this.Data["json"] = &data
	this.ServeJSON()
}

//添加文章 数据校验  路由 /api/article/add
func (this *ApiController) ArticleAdd() {
	data := &models.Article{}
	info := &ResultData{}
	if err := this.ParseForm(data); err != nil {
		info = &ResultData{Error:1,Title:"结果:",Msg:"接收表单数据出错！"}
	}else{
		data.Renew = tools.Int64ToString(time.Now().Unix())
		err := models.AddArticle(data)
		if err != nil{
			info = &ResultData{Error:1,Title:"结果:",Msg:"数据库操作出错！"}
		}else{
			info = &ResultData{Error:0,Title:"结果:",Msg:"发布成功！"}
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
		info = &ResultData{Error:1,Title:"结果:",Msg:"接收表单数据出错！"}
	}else{
		data.Id = tools.StringToInt(id)
		data.Renew = tools.Int64ToString(time.Now().Unix())
		err := models.UpdateArticle(data)
		if err != nil{
			info = &ResultData{Error:1,Title:"结果:",Msg:"数据库操作出错！"}
		}else{
			info = &ResultData{Error:0,Title:"结果:",Msg:"修改成功！"}
		}
	}
	this.Data["json"] = info
	this.ServeJSON()
}

// 文章删除  路由 /api/article/delete
func (this *ApiController) ArticleDelete() {
	info := &ResultData{}
	id,_ := strconv.Atoi(this.GetString("id"))
	err := models.DeleteArticle(id)
	if err != nil{
		info = &ResultData{Error:1,Title:"结果:",Msg:"数据库操作出错！"}
	}else{
		info = &ResultData{Error:0,Title:"结果:",Msg:"删除成功！"}
	}
	this.Data["json"] = info
	this.ServeJSON()
}

// 添加分类 路由 api/category/add
func (this *ApiController) CategoryAdd() {
	name := this.GetString("name")
	key := this.GetString("key")
	info := &models.Taxonomy{Name:name,Key:key}
	err := models.AddCategory(info)
	var data *ResultData
	if err != nil {
		data = &ResultData{Error:1,Title:"结果:",Msg:"添加失败！"}
	}else{
		data = &ResultData{Error:0,Title:"结果:",Msg:"添加成功！"}
	}
	this.Data["json"] = data
	this.ServeJSON()
}

// 网站设置页面  路由  /api/site/config
func (this *ApiController) SiteConfig() {
	submit := this.GetString("submit")
	info := &ResultData{}
	data := &models.ConfigOption{}
	if submit == "user" {
		config := &models.Config{Option:"Author",Value:this.GetString("username")}
		err := models.SiteConfig(config)
		config = &models.Config{Option:"Password",Value:this.GetString("password")}
		err1 := models.SiteConfig(config)
		if err != nil || err1 !=nil {
			info = &ResultData{Error:1,Title:"结果:",Msg:"出现未知错误！"}
		}else {
			info = &ResultData{Error:0,Title:"结果:",Msg:"信息重置成功！"}
		}
	}else{
		if err := this.ParseForm(data); err != nil {
			info = &ResultData{Error:1,Title:"结果:",Msg:"接收表单数据出错！"}
		}else{
			t := reflect.TypeOf(*data)	//类型
			v := reflect.ValueOf(*data)	//值
			for i := 0; i < t.NumField(); i++ {
				config := &models.Config{Option:t.Field(i).Name,Value:v.Field(i).String()}
				err := models.SiteConfig(config)
				if err != nil {
					info = &ResultData{Error:1,Title:"结果:",Msg:"出现未知错误！"}
				}else {
					info = &ResultData{Error:0,Title:"结果:",Msg:"信息重置成功！"}
				}
			}
		}
	}
	this.Data["json"] = info
	this.ServeJSON()
}