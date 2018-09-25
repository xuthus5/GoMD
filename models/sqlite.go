package models

import (
	"GoMD/tools"
	"github.com/astaxie/beego/orm"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"time"
)

var dbc orm.Ormer
var dbx *sqlx.DB

func init() {
	// 注册驱动
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	// 注册默认数据库
	// 备注：此处第一个参数必须设置为“default”（因为我现在只有一个数据库），否则编译报错说：必须有一个注册DB的别名为 default
	orm.RegisterDataBase("default", "sqlite3", "data.db")
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Config),new(Article),new(Taxonomy),new(Link),new(Message),new(Notice))
	// 开启 orm 调试模式：开发过程中建议打开，release时需要关闭
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)

	dbc = orm.NewOrm()
	dbc.Using("default")
	//安装初始化
	Initialization()
	//sqlx
	dbx,_ = sqlx.Open("sqlite3","data.db")
}

//安装初始化
func Initialization(){
	if _, err := os.Stat("install.lock"); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			//配置表初始化
			dbc.Insert(&Config{Option: "WebTitle", Value: "GoMD的博客"})
			dbc.Insert(&Config{Option: "Author", Value: "admin"})
			dbc.Insert(&Config{Option: "Password", Value: "admin"})
			dbc.Insert(&Config{Option: "Theme", Value: "default"})
			dbc.Insert(&Config{Option: "CopyRight",Value:"GoMD"})
			dbc.Insert(&Config{Option: "PageSize",Value:"10"})
			//文章分类表初始化
			dbc.Insert(&Taxonomy{
				Name: "默认",
				Key:  "default",
			})
			//文章表初始化
			dbc.Insert(&Article{
				Title:   "你好！世界",
				Cid:     1,
				Tags:    "文章",
				Summary: "你好！这是系统为你生成的第一篇文章！",
				Content: "你好！这是系统为你生成的第一篇文章！",
				Author: "administrator",
				Renew:  tools.Int64ToString(time.Now().Unix()),
			})

			file, err := os.Create("install.lock")
			if err != nil {
				log.Fatal(err.Error())
			}
			file.Close()
		}
	}
}