package main

import (
	"GoMD/models"
	"GoMD/tools"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// beego 日志配置结构体
type LoggerConfig struct {
	FileName            string   `json:"filename"`
	Level               int      `json:"level"`    // 日志保存的时候的级别，默认是 Trace 级别
	Maxlines            int      `json:"maxlines"` // 每个文件保存的最大行数，默认值 1000000
	Maxsize             int      `json:"maxsize"`  // 每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB
	Daily               bool     `json:"daily"`    // 是否按照每天 logrotate，默认是 true
	Maxdays             int      `json:"maxdays"`  // 文件最多保存多少天，默认保存 7 天
	Rotate              bool     `json:"rotate"`   // 是否开启 logrotate，默认是 true
	Perm                string   `json:"perm"`     // 日志文件权限
	RotatePerm          string   `json:"rotateperm"`
	EnableFuncCallDepth bool     `json:"-"` // 输出文件名和行号
	LogFuncCallDepth    int      `json:"-"` // 函数调用层级
	Separate            []string `json:"separate"`	//需要单独写入文件的日志级别,设置后命名类似 test.error.log
}

//日志运行级别检测变量
func logSetting() {
	//获取运行方式
	loglevel := beego.AppConfig.String("runmode")
	if loglevel == "prod" {
		//处于生产模式,检测日志运行类型
		logtype := beego.AppConfig.String("prod::logtype")
		logname := beego.AppConfig.String("prod::logname")
		separate := []string{"notice", "error"}
		var logCfg = LoggerConfig{
			FileName:            "logs/"+logname,
			Level:               7,
			EnableFuncCallDepth: true,
			LogFuncCallDepth:    3,
			RotatePerm:          "777",
			Perm:                "777",
			Separate:			 separate,
		}
		if logtype == "file" {
			err := tools.CheckAndDirCreate("logs")
			if err == nil {
				cfg,_ := json.Marshal(&logCfg)
				beego.SetLogger(logs.AdapterMultiFile,string(cfg))
			}
		}
	}
}

//自定义模板函数
func customTemplateFunction() {
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
}

//自定义状态页面
func statusCode() {
	beego.ErrorHandler("404", PageNotFound)
}

//自定义静态资源路径
func staticSourcePath() {
	beego.SetStaticPath("/file", "file")
}
