package models

import (
	"GoMD/tools"
	"github.com/astaxie/beego/orm"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/satori/go.uuid"
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
	orm.RegisterModel(new(Config), new(Article), new(Comment), new(Category), new(Link), new(Message), new(Notice), new(Attachment))
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
	dbc.Insert(&Config{Option: "Theme", Value: "QuietV1"})
	dbc.Insert(&Config{Option: "CopyRight", Value: "GoMD"})
	dbc.Insert(&Config{Option: "PageSize", Value: "10"})
	dbc.Insert(&Config{Option: "SecondaryTitle", Value: "金鳞岂是池中物，一遇风雨便化龙"})
	dbc.Insert(&Config{Option: "UserImageUrl", Value: "/static/common/images/user-head-image.jpeg"})
	dbc.Insert(&Config{Option: "LogoUrl", Value: "/static/common/images/user-head-image.jpeg"})
	dbc.Insert(&Config{Option: "UserEmail", Value: "admin@admin.com"})
	dbc.Insert(&Config{Option: "UserGithub", Value: "http://github.com"})
	//分类表初始化
	dbc.Insert(&Category{
		Name:        "默认",
		Key:         "default",
		Description: "默认的分类",
	})
	//文章表初始化
	dbc.Insert(&Article{
		Title:   "你好！世界",
		Uuid:    uuid.Must(uuid.NewV4()).String(),
		Cid:     1,
		Tags:    "文章",
		Summary: "你好！这是系统为你生成的第一篇文章！",
		Content: "你好！这是系统为你生成的第一篇文章！",
		Author:  "admin",
		Renew:   tools.Int64ToString(time.Now().Unix()),
	})
	//评论表初始化
	dbc.Insert(&Comment{
		Aid:     1,
		Content: "这是系统为你生成的一条评论。",
		Name:    "GoMD",
		Link:    "/",
		Email:   "1397190480@qq.com",
		Date:    tools.Int64ToString(time.Now().Unix()),
	})
	//公告表初始化
	dbc.Insert(&Notice{
		Content: "你好！欢迎使用GoMD",
		Date:    tools.Int64ToString(time.Now().Unix()),
	})
	//附件表初始化
	dbc.Insert(&Attachment{
		Name:    "user-head-image.jpeg",
		Path:    "file/user-head-image.jpeg",
		Created: tools.Int64ToString(time.Now().Unix()),
	})
	//友情链接表初始化
	dbc.Insert(&Link{
		Name:        "GoMD",
		Url:         "https://gitee.com/xuthus5/GoMD",
		Description: "GoMD源代码托管地址",
	})
}
