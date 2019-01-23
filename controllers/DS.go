package controllers

/**************
控制器中使用的数据结构
**************/

/*
返回操作结果
	0 没有错误
	1 有错误
*/
type ResultData struct {
	Error int64                     //错误代码
	Title string                    //标题
	Msg   string                    //信息
	Data  interface{} `json:"data"` //数据
}

/*
 echo category information
 */

type CategoryData struct {
	IsNil 	bool	//check this category is null ?
	Data    interface{} 	//the article include in this category content
}