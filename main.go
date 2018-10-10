package main

import (
	_ "GoMD/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.AddFuncMap("tags", Tags)
	beego.AddFuncMap("calc", Calc)
	beego.AddFuncMap("markdown", MarkDown)
	beego.AddFuncMap("time", YMD)
	beego.AddFuncMap("sc", SiteConfig)
	beego.AddFuncMap("tableNum", TableNumber)
	beego.AddFuncMap("notice", GetNotice)
	beego.ErrorHandler("404",PageNotFound)
	beego.Run()
}

