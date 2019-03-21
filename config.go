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
	EnableFuncCallDepth bool     `json:"-"`        // 输出文件名和行号
	LogFuncCallDepth    int      `json:"-"`        // 函数调用层级
	Separate            []string `json:"separate"` //需要单独写入文件的日志级别,设置后命名类似 test.error.log
}

//
type ChangeLogData struct {
	Email   string `json:"Email"`
	Name    string `json:"Name"`
	Version []struct {
		Commit    string `json:"Commit"`
		Committer string `json:"Committer"`
		Date      string `json:"Date"`
		Tag       string `json:"Tag"`
		Title     string `json:"Title"`
		Type      string `json:"Type"`
		Level     string `json:"Level"`
		Version   string `json:"Version"`
	} `json:"Version"`
	Website string `json:"Website"`
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
			FileName:            "logs/" + logname,
			Level:               7,
			EnableFuncCallDepth: true,
			LogFuncCallDepth:    3,
			RotatePerm:          "777",
			Perm:                "777",
			Separate:            separate,
		}
		if logtype == "file" {
			err := tools.CheckAndDirCreate("logs")
			if err == nil {
				cfg, _ := json.Marshal(&logCfg)
				_ = beego.SetLogger(logs.AdapterMultiFile, string(cfg))
			}
		}
	}
}

//自定义模板函数
func customTemplateFunction() {
	//时间处理相关
	_ = beego.AddFuncMap("tsc", TimeStampConversion) //将时间戳转成正常时间
	_ = beego.AddFuncMap("ed", EnumerateDate)        //返回单独的年月日
	_ = beego.AddFuncMap("dt", DivisionTime)         //切割时间为 ymd : hms

	//运算相关
	_ = beego.AddFuncMap("calc", Calc) //加减计算

	//数据处理输出
	_ = beego.AddFuncMap("tags", Tags)         //拆分标签
	_ = beego.AddFuncMap("markdown", MarkDown) //将markdown输出为html

	//数据关系处理
	_ = beego.AddFuncMap("gafc", GetArticleFromCommentID)
	_ = beego.AddFuncMap("changelog", ChangeLog)

	//数据获取
	_ = beego.AddFuncMap("sc", SiteConfig)                //调取网站配置 直接抽调数据库配置字段
	_ = beego.AddFuncMap("tn", TableNumber)               //获取表中的数据，用于统计数据表中数据的条数
	_ = beego.AddFuncMap("na", models.GetLimitNewArticle) //获取指定条数的最近更新的文章
	_ = beego.AddFuncMap("nc", models.GetLimitNewComment) //获取指定条数的最近的评论
	_ = beego.AddFuncMap("notice", GetNotice)             //网站的公告
	_ = beego.AddFuncMap("category", GetCategory)         //文章的分类信息
	_ = beego.AddFuncMap("cn", GetAOfCategoryNumber)      //获取分类下的文章数量
	_ = beego.AddFuncMap("comment", GetCommentNumber)     //获取分类下的评论数量
	_ = beego.AddFuncMap("gavatar", GetGravatar)          //获取评论者gavatar头像
	_ = beego.AddFuncMap("pna", PreOrNextAriticle)        //获取上一篇或者下一篇文章
	_ = beego.AddFuncMap("itu", IdToName)                 //根据id返回name
	_ = beego.AddFuncMap("cl", models.CategoryList)       //返回分类列表
	_ = beego.AddFuncMap("link", models.GetAllLink)       //返回链接列表
	_ = beego.AddFuncMap("menu", models.GetAllMenu)       //返回菜单栏
	_ = beego.AddFuncMap("gpb", models.GetPropertyByID)   //通过提供ID和属性字段,返回字段内容
}

//自定义状态页面
func statusCode() {
	beego.ErrorHandler("404", PageNotFound)
}

//自定义静态资源路径
func staticSourcePath() {
	beego.SetStaticPath("/file", "file")
}
