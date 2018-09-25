package main

import (
	"GoMD/models"
	"gopkg.in/russross/blackfriday.v2"
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

/* 获取网站基本信息 */

/* 返回表中的数据条数 用于数据统计 */
func TableNumber(table string) int64{
	count,_ := models.TableNumber(table)
	return count
}