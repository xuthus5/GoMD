package main

import (
	_ "GoMD/routers"
	"github.com/astaxie/beego"
	"html/template"
	"net/http"
)

//自定义404报错
func PageNotFound(rw http.ResponseWriter, r *http.Request){
	t,_:= template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath+"/default/404.html")
	data := make(map[string]interface{})
	data["code"] = "404"
	data["title"] = "页面被吃掉了！"
	t.Execute(rw, data)
}

func main() {
	// 运行时
	beego.AddFuncMap("tags", Tags)
	beego.AddFuncMap("calc", Calc)
	beego.AddFuncMap("markdown", MarkDown)
	beego.AddFuncMap("time", YMD)
	beego.AddFuncMap("tableNum", TableNumber)
	beego.ErrorHandler("404",PageNotFound)
	beego.Run()
}

