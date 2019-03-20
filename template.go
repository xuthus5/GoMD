package main

import (
	"GoMD/models"
	"GoMD/tools"
	"encoding/json"
	"github.com/astaxie/beego"
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/* ---------------------------------
功能： 模板函数文件
------------------------------------*/

/*******************
时间处理相关
********************/

/* 时间转换 传入时间戳字符串 是否精准 */
func TimeStampConversion(timestamp string, exact bool) string {
	realTime, _ := strconv.ParseInt(timestamp, 10, 64)
	if exact == false {
		return time.Unix(realTime, 0).Format("2006-01-02")
	} else {
		return time.Unix(realTime, 0).Format("2006-01-02 15:04:05")
	}
}

/* 得到单独的年月日 */
func EnumerateDate(method string) string {
	year, mouth, day := tools.EnumerateDate()
	if method == "year" {
		return year
	} else if method == "mouth" {
		return mouth
	} else if method == "day" {
		return day
	} else {
		return ""
	}
}

/*******************
运算相关
********************/
/* 分页处理 */
func Calc(x, y int64, option string) int64 {
	switch option {
	case "+":
		return x + y
	case "-":
		return x - y
	default:
		return x + y
	}
}

/*******************
数据处理输出
********************/
/* 文章页面标签显示 */
func Tags(tags string) []string {
	array := strings.Split(tags, ",")
	return array
}

/* markdown转换 */
func MarkDown(content string) string {
	output := blackfriday.Run([]byte(content), blackfriday.WithExtensions(blackfriday.NoEmptyLineBeforeBlock))
	return string(output)
}

/*******************
数据关系处理
********************/
/* 由评论ID得到文章信息 */
func GetArticleFromCommentID(id int, fields string) string {
	data := models.GetOneArticle(strconv.Itoa(id), "id")
	article := *data
	if fields == "title" {
		return article[0].Title
	}
	return strconv.Itoa(id)
}

/*******************
数据获取
********************/
/* 直接调取数据库配置表-网站基本信息 */
func SiteConfig(info string) string {
	return models.GetOneConfig(info)
}

/* 返回表中的数据条数 用于数据统计 */
func TableNumber(table string) int64 {
	count, _ := models.TableNumber(table)
	return count
}

/* 返回分类下的文章数量 */
func GetAOfCategoryNumber(cid int) int64 {
	count, _ := models.GetNumFromACategory(cid)
	return count
}

/* 获取公告表数据 */
func GetNotice() string {
	return models.GetOneNotice()
}

/* 根据分类cid 返回对应的分类名称 */
func GetCategory(id int, method string) (value string) {
	temp := models.GetOneCategoryInfo(strconv.Itoa(id), "id")
	category := *temp
	if method == "name" {
		value = category[0].Name
	} else {
		value = category[0].Key
	}
	return
}

/* 返回评论的头像地址 */
func GetGravatar(email string) string {
	result := tools.StringToMd5(email, 0)
	return "http://cn.gravatar.com/avatar/" + result + "&d=identicon"
}

/* 获取当前文章的评论数量 */
func GetCommentNumber(id int) int64 {
	num, _ := models.GetCommentNumFromArticle(id)
	return num
}

/* 获得当前文章的上下篇文章 */
func PreOrNextAriticle(id int, method string) *map[string]string {
	if method == "pre" {
		return models.GetPreOrNextArticle(id, "pre")
	} else {
		return models.GetPreOrNextArticle(id, "next")
	}
}

/* 根据id返回name */
func IdToName(id int) string {
	ids := int64(id)
	return models.GetNameById(ids)
}

//自定义404报错
func PageNotFound(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("error.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/admin/error.html")
	data := make(map[string]interface{})
	data["code"] = "404"
	data["title"] = "页面被吃掉了！"
	_ = t.Execute(rw, data)
}

/* 切割时间为 ymd : hms */
func DivisionTime(stamp string) map[string]string {
	ymd, hms := tools.DivisionTime(stamp)
	getTime := make(map[string]string)
	getTime["ymd"] = ymd
	getTime["hms"] = hms
	return getTime
}

//官方最新日志获取
func ChangeLog() ChangeLogData {
	data := ChangeLogData{}
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	resp, err := http.Get("https://gitee.com/xuthus5/GoMD/raw/master/CHANGELOG.json")
	if err != nil {
		beego.Error("获取更新信息失败")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Error("读取更新信息失败！")
	}
	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(body, &data)
	if err != nil {
		beego.Error("解析json数据失败")
	}
	return data
}

//获取菜单栏
func MenuList() *[]models.Link {
	return models.GetAllMenu()
}
