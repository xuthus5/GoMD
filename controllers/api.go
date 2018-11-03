package controllers

import (
	"GoMD/models"
	"GoMD/tools"
	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"log"
	"reflect"
	"strconv"
	"time"
)

type ApiController struct {
	beego.Controller
}

// 返回json列表 数据格式
type JsonData struct {
	Code  int         `json:"code"`  //错误代码
	Count int         `json:"count"` // 数据数量
	Msg   string      `json:"msg"`   //输出信息
	Data  interface{} `json:"data"`  //数据
}

//返回后台文章列表 json数据类型返回 路由 /api/article/list
func (this *ApiController) ArticleList() {
	this.Data["json"] = &JsonData{Code: 0, Count: 100, Msg: "", Data: models.GetArticleJson()}
	this.ServeJSON()
}

//返回后台分类列表 json数据类型返回 路由 /api/article/category/list
func (this *ApiController) CategoryList() {
	this.Data["json"] = &JsonData{Code: 0, Count: 100, Msg: "", Data: models.GetCategoryJson()}
	this.ServeJSON()
}

//添加文章 数据校验  路由 /api/article/add
func (this *ApiController) ArticleAdd() {
	data := &models.Article{}
	info := &ResultData{}
	if err := this.ParseForm(data); err != nil {
		info = &ResultData{Error: 1, Title: "失败:", Msg: "接收表单数据出错！"}
	} else {
		data.Renew = tools.Int64ToString(time.Now().Unix())
		data.Uuid = uuid.Must(uuid.NewV4()).String()
		idstr, err := models.AddArticle(data)
		if err != nil {
			info = &ResultData{Error: 1, Title: "失败:", Msg: "数据库操作出错！"}
		} else {
			info = &ResultData{Error: 0, Title: "成功:", Msg: "发布成功！", Data: idstr}
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
		info = &ResultData{Error: 1, Title: "失败:", Msg: "接收表单数据出错！"}
	} else {
		data.Id = tools.StringToInt(id)
		data.Renew = tools.Int64ToString(time.Now().Unix())
		data.Uuid = uuid.Must(uuid.NewV4()).String()
		err = models.UpdateArticle(data)
		if err != nil {
			info = &ResultData{Error: 1, Title: "失败:", Msg: "数据库操作出错！"}
		} else {
			info = &ResultData{Error: 0, Title: "成功:", Msg: "修改成功！",Data: data.Uuid}
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
		info = &ResultData{Error: 1, Title: "失败:", Msg: "数据库操作出错！"}
	} else {
		info = &ResultData{Error: 0, Title: "成功:", Msg: "删除成功！"}
	}
	this.Data["json"] = info
	this.ServeJSON()
}

// 添加分类 路由 /api/article/category/add
func (this *ApiController) CategoryAdd() {
	name := this.GetString("name")
	key := this.GetString("key")
	description := this.GetString("description")
	info := &models.Category{Name: name, Key: key, Description: description}
	err := models.AddCategory(info)
	var data *ResultData
	if err != nil {
		data = &ResultData{Error: 1, Title: "失败:", Msg: "添加失败！"}
	} else {
		data = &ResultData{Error: 0, Title: "成功:", Msg: "添加成功！"}
	}
	this.Data["json"] = data
	this.ServeJSON()
}

// 分类删除 路由 /api/article/category/delete
func (this *ApiController) CategoryDelete() {
	info := &ResultData{}
	//先判断分类数目 为1时，不允许删除
	count, _ := models.TableNumber("category")
	if count == 1 {
		info = &ResultData{Error: 1, Title: "失败:", Msg: "必须保留一个分类！"}
	} else {
		id, _ := strconv.Atoi(this.GetString("id"))
		err := models.DeleteCategory(id)
		if err != nil {
			info = &ResultData{Error: 1, Title: "失败:", Msg: "数据库操作出错！"}
		} else {
			info = &ResultData{Error: 0, Title: "成功:", Msg: "删除成功！"}
		}
	}
	this.Data["json"] = info
	this.ServeJSON()
}

// 分类修改 路由 /api/article/category/update
func (this *ApiController) CategoryUpdate() {
	id := this.GetString("id")
	data := &models.Category{}
	info := &ResultData{}
	if err := this.ParseForm(data); err != nil {
		info = &ResultData{Error: 1, Title: "失败:", Msg: "接收表单数据出错！"}
	} else {
		data.Id = tools.StringToInt(id)
		err := models.UpdateCategory(data)
		if err != nil {
			info = &ResultData{Error: 1, Title: "失败:", Msg: "数据库操作出错！"}
		} else {
			info = &ResultData{Error: 0, Title: "成功:", Msg: "修改成功！"}
		}
	}
	this.Data["json"] = info
	this.ServeJSON()
}

// 网站设置页面  路由  /api/site/config
func (this *ApiController) SiteConfig() {
	// 判断提交类型 user为用户信息表单  site为网站配置表单
	submit := this.GetString("submit")
	info := &ResultData{}
	var data interface{}
	if submit == "user" {
		data = &models.UserConfigOption{}
	} else {
		data = &models.SiteConfigOption{}
	}
	if err := this.ParseForm(data); err != nil {
		log.Println(data)
		info = &ResultData{Error: 1, Title: "失败:", Msg: "接收表单数据出错！"}
	} else {
		t := reflect.TypeOf(data).Elem()  //类型
		v := reflect.ValueOf(data).Elem() //值
		for i := 0; i < t.NumField(); i++ {
			config := &models.Config{Option: t.Field(i).Name, Value: v.Field(i).String()}
			err := models.SiteConfig(config)
			if err != nil {
				info = &ResultData{Error: 1, Title: "失败:", Msg: "出现未知错误！"}
			} else {
				info = &ResultData{Error: 0, Title: "成功:", Msg: "信息重置成功！"}
			}
		}
	}
	this.Data["json"] = info
	this.ServeJSON()
}

// 公告数据接受 路由 /api/notice
func (this *ApiController) Notice() {
	notice := &models.Notice{Content: this.GetString("content"), Url: this.GetString("url"), Date: tools.Int64ToString(time.Now().Unix())}
	info := &ResultData{}
	if models.AddNotice(notice) != nil {
		info = &ResultData{Error: 1, Title: "失败:", Msg: "发布失败！"}
	} else {
		info = &ResultData{Error: 0, Title: "成功:", Msg: "发布成功！"}
	}
	this.Data["json"] = info
	this.ServeJSON()
}

// 公告列表 路由 /api/notice/list
func (this *ApiController) NoticeList() {
	this.Data["json"] = &JsonData{Code: 0, Count: 100, Msg: "", Data: models.GetNoticeJson()}
	this.ServeJSON()
}

// 文件上传api 路由 /api/file/upload 返回一个包含文件存储信息的json数据
func (this *ApiController) FileUpload() {
	info := &ResultData{Error: 1, Title: "失败:", Msg: "上传失败！"}
	f, h, err := this.GetFile("file")
	if err != nil {
		log.Fatal("error: ", err)
	}
	defer f.Close()
	//获取当前年月日
	year, month, _ := tools.EnumerateDate()
	savePath := "file/" + year + "/" + month + "/"
	//创建存储目录
	tools.DirCreate(savePath)
	//重命名文件名称
	tempFileName := tools.StringToMd5(h.Filename, 5)
	suffix := tools.GetFileSuffix(h.Filename)
	saveName := tempFileName + suffix
	// 保存位置
	err = this.SaveToFile("file", savePath+saveName)
	//写入数据库
	if err == nil {
		//写入数据库
		data := &models.Attachment{Name: saveName, Path: savePath + saveName, Created: tools.Int64ToString(time.Now().Unix())}
		id, code := models.FileSave(data)
		if code != nil {
			info = &ResultData{Error: 1, Title: "结果:", Msg: "上传失败！"}
		} else {
			info = &ResultData{Error: 0, Title: "结果:", Msg: "上传成功！", Data: models.FileInfo(id)}
		}
	}
	this.Data["json"] = info
	this.ServeJSON()
}

//文件列表api 路由 /api/file/list
func (this *ApiController) FileList() {
	this.Data["json"] = &JsonData{Code: 0, Count: 100, Msg: "", Data: models.GetFileJson()}
	this.ServeJSON()
}

// 文件删除 路由 /api/file/delete
func (this *ApiController) FileDelete() {
	info := &ResultData{}
	id, _ := strconv.Atoi(this.GetString("id"))
	//数据库文件删除
	path, err := models.FileDelete(id)
	if err != nil {
		info = &ResultData{Error: 1, Title: "失败:", Msg: "数据库操作出错！"}
	} else {
		info = &ResultData{Error: 0, Title: "成功:", Msg: "删除成功！"}
	}
	//本地文件删除
	tools.FileRemove(path)
	this.Data["json"] = info
	this.ServeJSON()
}

// 发布评论 路由 /api/comment/add
func (this *ApiController) CommentAdd() {
	data := &models.Comment{}
	info := &ResultData{}
	if err := this.ParseForm(data); err != nil {
		info = &ResultData{Error: 1, Title: "失败:", Msg: "接收表单数据出错！"}
	} else {
		data.Date = tools.Int64ToString(time.Now().Unix())
		err := models.AddComment(data)
		if err != nil {
			info = &ResultData{Error: 1, Title: "失败:", Msg: "数据库操作出错！"}
		} else {
			info = &ResultData{Error: 0, Title: "成功:", Msg: "发布成功！"}
		}
	}
	this.Data["json"] = info
	this.ServeJSON()
}
