package models

import (
	"GoMD/tools"
	"github.com/astaxie/beego/orm"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

var dbc orm.Ormer
var dbx *sqlx.DB

func init() {
	// 注册驱动
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	// 注册默认数据库
	orm.RegisterDataBase("default", "sqlite3", "data.db")
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Config), new(Article), new(Taxonomy), new(Link), new(Message), new(Notice), new(Source))
	// 开启 orm 调试模式：开发过程中建议打开，release时需要关闭
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)

	dbc = orm.NewOrm()
	dbc.Using("default")

	//检测安装初始化
	isInstall := dbc.Read(&Config{Option: "IsInstall", Value: "yes"}, "Option", "Value")
	if isInstall != nil {
		Initialization()
	}
	//sqlx
	dbx, _ = sqlx.Open("sqlite3", "data.db")
}

//安装初始化
func Initialization() {
	//配置表初始化
	dbc.Insert(&Config{Option: "IsInstall", Value: "yes"})
	dbc.Insert(&Config{Option: "WebTitle", Value: "GoMD的博客"})
	dbc.Insert(&Config{Option: "Author", Value: "admin"})
	dbc.Insert(&Config{Option: "Password", Value: "admin"})
	dbc.Insert(&Config{Option: "Theme", Value: "default"})
	dbc.Insert(&Config{Option: "CopyRight", Value: "GoMD"})
	dbc.Insert(&Config{Option: "PageSize", Value: "10"})
	//文章分类表初始化
	dbc.Insert(&Taxonomy{
		Name: "默认",
		Key:  "default",
		Description:"默认的分类",
	})
	//文章表初始化
	dbc.Insert(&Article{
		Title:   "你好！世界",
		Cid:     1,
		Tags:    "文章",
		Summary: "你好！这是系统为你生成的第一篇文章！",
		Content: "你好！这是系统为你生成的第一篇文章！",
		Author:  "admin",
		Renew:   tools.Int64ToString(time.Now().Unix()),
	})
	//公告表初始化
	dbc.Insert(&Notice{
		Content: "你好！欢迎使用GoMD",
		Data:    tools.Int64ToString(time.Now().Unix()),
	})
}
