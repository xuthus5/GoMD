package controllers

import (
	"GoMD/models"
	"GoMD/tools"
	"github.com/Lofanmi/pinyin-golang/pinyin"
	"github.com/astaxie/beego"
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
	Msg   string      `json:"msg"`   //输出信息
	Count int         `json:"total"` // 数据数量
	Data  interface{} `json:"data"`  //数据
}

/************************

有关文章的API

*************************/

//返回后台文章列表 json数据类型返回 路由 /api/article/list
func (this *ApiController) ArticleList() {
	data := models.GetArticleJson()
	this.Data["json"] = &JsonData{Code: 0, Count: len(*data), Msg: "", Data: data}
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
	dict := pinyin.NewDict()
	if err := this.ParseForm(data); err != nil {
		info = &ResultData{Error: 1, Title: "失败:", Msg: "接收表单数据出错！"}
	} else {
		data.Renew = tools.Int64ToString(time.Now().Unix())
		data.Uuid = dict.Convert(data.Title, "-").None()
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
	dict := pinyin.NewDict()
	if err := this.ParseForm(data); err != nil {
		info = &ResultData{Error: 1, Title: "失败:", Msg: "接收表单数据出错！"}
	} else {
		data.Id = tools.StringToInt(id)
		data.Renew = tools.Int64ToString(time.Now().Unix())
		data.Uuid = dict.Convert(data.Title, "-").None()
		err = models.UpdateArticle(data)
		if err != nil {
			info = &ResultData{Error: 1, Title: "失败:", Msg: "数据库操作出错！"}
		} else {
			info = &ResultData{Error: 0, Title: "成功:", Msg: "修改成功！", Data: data.Uuid}
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
		err := models.DeleteCommentFromArticle(id)
		if err == nil {
			info = &ResultData{Error: 0, Title: "成功:", Msg: "删除成功！"}
		} else {
			info = &ResultData{Error: 1, Title: "失败:", Msg: "数据库操作出错！"}
		}
	}
	this.Data["json"] = info
	this.ServeJSON()
}

/************************

有关页面的API

*************************/

//添加页面 路由 /api/page/add
func (this *ApiController) PageAdd() {
	data := &models.Article{}
	info := &ResultData{}
	dict := pinyin.NewDict()
	if err := this.ParseForm(data); err != nil {
		info = &ResultData{Error: 1, Title: "失败:", Msg: "接收表单数据出错！"}
	} else {
		data.Renew = tools.Int64ToString(time.Now().Unix())
		data.Type = 1
		if data.Uuid == "" {
			data.Uuid = dict.Convert(data.Title, "-").None()
		}
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

// 页面更新 数据校验  路由 /api/page/update
func (this *ApiController) PageUpdate() {
	id := this.GetString("id")
	data := &models.Article{}
	info := &ResultData{}
	dict := pinyin.NewDict()
	if err := this.ParseForm(data); err != nil {
		info = &ResultData{Error: 1, Title: "失败:", Msg: "接收表单数据出错！"}
	} else {
		data.Id = tools.StringToInt(id)
		data.Type = 1
		data.Renew = tools.Int64ToString(time.Now().Unix())
		if data.Uuid == "" {
			data.Uuid = dict.Convert(data.Title, "-").None()
		}
		err = models.UpdateArticle(data)
		if err != nil {
			info = &ResultData{Error: 1, Title: "失败:", Msg: "数据库操作出错！"}
		} else {
			info = &ResultData{Error: 0, Title: "成功:", Msg: "修改成功！", Data: data.Uuid}
		}
	}
	this.Data["json"] = info
	this.ServeJSON()
}

//页面删除  路由 /api/page/delete
func (this *ApiController) PageDelete() {
	info := &ResultData{}
	id, _ := strconv.Atoi(this.GetString("id"))
	err := models.DeleteArticle(id)
	if err != nil {
		info = &ResultData{Error: 1, Title: "失败:", Msg: "数据库操作出错！"}
	} else {
		err := models.DeleteCommentFromArticle(id)
		if err == nil {
			info = &ResultData{Error: 0, Title: "成功:", Msg: "删除成功！"}
		} else {
			info = &ResultData{Error: 1, Title: "失败:", Msg: "数据库操作出错！"}
		}
	}
	this.Data["json"] = info
	this.ServeJSON()
}

//返回后台页面列表 json数据类型返回 路由 /api/page/list
func (this *ApiController) PageList() {
	data := models.GetPageJson()
	this.Data["json"] = &JsonData{Code: 0, Count: len(*data), Msg: "", Data: data}
	this.ServeJSON()
}
/************************

有关分类的API

*************************/

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

/************************

有关网站配置的API

*************************/

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

/************************

有关公告的API

*************************/

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

/************************

有关文件处理的API

*************************/

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
	tools.CheckAndDirCreate(savePath)
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

/************************

有关评论的API

*************************/

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

/************************

有关链接的API

*************************/

// 添加链接 路由 /api/link/add
func (this *ApiController) LinkAdd() {
	name := this.GetString("name")
	url := this.GetString("url")
	description := this.GetString("description")
	info := &models.Link{Name: name, Url: url, Description: description}
	err := models.AddLink(info)
	var data *ResultData
	if err != nil {
		data = &ResultData{Error: 1, Title: "失败:", Msg: "添加失败！"}
	} else {
		data = &ResultData{Error: 0, Title: "成功:", Msg: "添加成功！"}
	}
	this.Data["json"] = data
	this.ServeJSON()
}

// 删除链接 路由 /api/link/delete
func (this *ApiController) LinkDelete() {
	info := &ResultData{}

	id, _ := strconv.Atoi(this.GetString("id"))
	err := models.DeleteLink(id)
	if err != nil {
		info = &ResultData{Error: 1, Title: "失败:", Msg: "数据库操作出错！"}
	} else {
		info = &ResultData{Error: 0, Title: "成功:", Msg: "删除成功！"}
	}

	this.Data["json"] = info
	this.ServeJSON()
}

//返回后台链接列表 json数据类型返回 路由 /api/link/list
func (this *ApiController) LinkList() {
	this.Data["json"] = &JsonData{Code: 0, Count: 100, Msg: "", Data: models.GetAllLink()}
	this.ServeJSON()
}

// 链接修改 路由 /api/link/update
func (this *ApiController) LinkUpdate() {
	id := this.GetString("id")
	data := &models.Link{}
	info := &ResultData{}
	if err := this.ParseForm(data); err != nil {
		info = &ResultData{Error: 1, Title: "失败:", Msg: "接收表单数据出错！"}
	} else {
		data.Id = tools.StringToInt(id)
		err := models.UpdateLink(data)
		if err != nil {
			info = &ResultData{Error: 1, Title: "失败:", Msg: "数据库操作出错！"}
		} else {
			info = &ResultData{Error: 0, Title: "成功:", Msg: "修改成功！"}
		}
	}
	this.Data["json"] = info
	this.ServeJSON()
}

/************************

有关菜单的API

*************************/

// 添加菜单栏节点 路由 /api/menu/add
func (this *ApiController) MenuNodeAdd() {
	name := this.GetString("name")
	url := this.GetString("url")
	description := this.GetString("description")
	info := &models.Link{Name: name, Url: url, Description: description, Type:1}
	err := models.AddLink(info)
	var data *ResultData
	if err != nil {
		data = &ResultData{Error: 1, Title: "失败:", Msg: "添加失败！"}
	} else {
		data = &ResultData{Error: 0, Title: "成功:", Msg: "添加成功！"}
	}
	this.Data["json"] = data
	this.ServeJSON()
}

// 删除菜单栏节点 路由 /api/menu/delete
func (this *ApiController) MenuNodeDelete() {
	info := &ResultData{}

	id, _ := strconv.Atoi(this.GetString("id"))
	err := models.DeleteLink(id)
	if err != nil {
		info = &ResultData{Error: 1, Title: "失败:", Msg: "数据库操作出错！"}
	} else {
		info = &ResultData{Error: 0, Title: "成功:", Msg: "删除成功！"}
	}

	this.Data["json"] = info
	this.ServeJSON()
}

//返回给后台一个菜单栏的列表 json数据类型返回 路由 /api/menu/list
func (this *ApiController) MenuList() {
	this.Data["json"] = &JsonData{Code: 0, Count: 100, Msg: "", Data: models.GetAllMenu()}
	this.ServeJSON()
}

// 菜单栏节点修改 路由 /api/menu/update
func (this *ApiController) MenuNodeUpdate() {
	id := this.GetString("id")
	data := &models.Link{}
	info := &ResultData{}
	if err := this.ParseForm(data); err != nil {
		info = &ResultData{Error: 1, Title: "失败:", Msg: "接收表单数据出错！"}
	} else {
		data.Id = tools.StringToInt(id)
		err := models.UpdateLink(data)
		if err != nil {
			info = &ResultData{Error: 1, Title: "失败:", Msg: "数据库操作出错！"}
		} else {
			info = &ResultData{Error: 0, Title: "成功:", Msg: "修改成功！"}
		}
	}
	this.Data["json"] = info
	this.ServeJSON()
}

/************************

有关评论的API

*************************/

// 待审核评论列表 路由 /api/comment/review
func (this *ApiController) ReviewComment()  {
	this.Data["json"] = &JsonData{Code: 0, Count: 100, Msg: "", Data: models.ReviewComment()}
	this.ServeJSON()
}

// 通过的评论 路由 /api/comment/adopt
func (this *ApiController) AdoptComment()  {
	this.Data["json"] = &JsonData{Code: 0, Count: 100, Msg: "", Data: models.AdoptComment()}
	this.ServeJSON()
}

// 删除评论 路由 /api/comment/delete
func (this *ApiController) DeleteComment() {
	info := &ResultData{}
	id, _ := strconv.Atoi(this.GetString("id"))
	err := models.CommentDelete(id)
	if err != nil {
		info = &ResultData{Error: 1, Title: "失败:", Msg: "数据库操作出错！"}
	} else {
		info = &ResultData{Error: 0, Title: "成功:", Msg: "删除成功！"}
	}
	this.Data["json"] = info
	this.ServeJSON()
}

// 通过评论 路由 /api/comment/update
func (this *ApiController) UpdateComment() {
	id := this.GetString("id")
	data := &models.Comment{}
	info := &ResultData{}
	data.Id = tools.StringToInt(id)
	data.Status = 1
	err := models.CommentUpdate(data)
	if err != nil {
		info = &ResultData{Error: 1, Title: "失败:", Msg: "数据库操作出错！"}
	} else {
		info = &ResultData{Error: 0, Title: "成功:", Msg: "修改成功！"}
	}
	this.Data["json"] = info
	this.ServeJSON()
}