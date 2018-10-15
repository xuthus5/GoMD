package main

import (
	"GoMD/models"
	"github.com/astaxie/beego"
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/* ---------------------------------
功能： 模板函数文件
------------------------------------*/

/* 文章页面标签显示 */
func Tags(tags string) []string {
	array := strings.Split(tags, ",")
	return array
}

/* 分页处理 */
func Calc(x,y int64,option string) int64{
	switch option {
	case "+":
		return x+y
	case "-":
		return x-y
	default:
		return x+y
	}
}

/* markdown转换 */
func MarkDown(content string) string{
	output := blackfriday.Run([]byte(content))
	return string(output)
}

/* 时间转换 传入时间戳字符串 是否精准 */
func YMD(timestamp string,exact bool) string{
	realTime,_ := strconv.ParseInt(timestamp, 10, 64)
	if exact == false{
		return time.Unix(realTime,0).Format("2006-01-02")
	}else{
		return time.Unix(realTime,0).Format("2006-01-02 15:04:05")
	}
}

/* 直接调取数据库配置表-网站基本信息 */
func SiteConfig(info string) string{
	return models.GetOneConfig(info)
}

/* 返回表中的数据条数 用于数据统计 */
func TableNumber(table string) int64{
	count,_ := models.TableNumber(table)
	return count
}

/* 返回分类下的文章数量 */
func GetAOfCategoryNumber(cid int) int64 {
	count,_ := models.GetNumFromACategory(cid)
	return count
}

/* 获取公告表数据 */
func GetNotice() string{
	return models.GetOneNotice()
}

/* 根据分类cid 返回对应的分类名称 */
func GetCategory(id int) string {
	ids := strconv.Itoa(id)
	temp := models.GetOneCategoryInfo(ids)
	category := *temp
	return category[0].Name
}

//自定义404报错
func PageNotFound(rw http.ResponseWriter, r *http.Request){
	t,_:= template.New("error.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath+"/common/error.html")
	data := make(map[string]interface{})
	data["code"] = "404"
	data["title"] = "页面被吃掉了！"
	t.Execute(rw, data)
}