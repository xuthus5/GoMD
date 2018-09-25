package models


/* ---------------------------

功能：database structure
数据库字段统一存放位置

------------------------------ */

/*
全局配置表
@ Option 	配置选项
@ Value 	配置结果
*/
type Config struct {
	Id			int
	Option 		string	`orm:"size(16)"`
	Value 		string	`orm:"size(32)"`
}

//文章表 存储完整的文章信息 映射文章数据表
type Article struct {
	Id 			int		`form:"-"`
	Cid 		int		`form:"cid"`
	Title 		string	`orm:"size(64)" form:"title"`
	Author 		string	`form:"author"`
	Tags 		string	`orm:"size(64)" form:"tags"`
	Renew 		string	`orm:"size(20)" form:"-"`
	Content 	string	`orm:"type(text)" form:"editormd-markdown-doc"`
	Summary		string	`form:"summary"`
	Status 		string	`form:"status"`
	Password	string	`form:"password"`
	Link 		string	`form:"-"`
}

// 链接表
type Link struct {
	Id 		int //自增唯一ID
	Url 	string //链接URL
	Name 	string //：链接标题
	Image	string //：链接图片
	Target  string //：链接打开方式
	Description  string //：链接描述
	Visible		string //：是否可见（Y/N）
}

//公告表
type Notice struct {
	ID  int //自增唯一ID
	Content string //内容
	data    string //发布时间
	Url     string //定向url
}

//留言
type Message struct {
	ID 		int
	Name 	string
	Email 	string
	Url		string
	Content 	string
	Data 	string
}

// taxonomy 分类表
type Taxonomy struct {
	Id 		int //：分类方法ID
	Name 	string 	//分类名称
	Key 	string  //分类标识
	Method  string //：分类方法(category/post_tag)
	Description string //：分类描述
	Parent  string //：所属父分类方法ID
}

/* 后台文章数据表结构
	id
	文章标题
	文章作者
	文章分类名称
	文章标签
	文章更新时间
*/
type DisplayArticle struct {
	Id 			int
	Title 		string
	Author 		string
	Name 		string
	Tags 		string
	Renew 		string
}

/*
网站配置结构体 持续添加
*/
type ConfigOption struct {
	WebTitle 	string	`form:"WebTitle"`
	Keywords 	string	`form:"keywords"`
	WebUrl 		string	`form:"WebUrl"`
	Description string	`form:"description"`
	Status 		string	`form:"status"`
	Theme 		string	`form:"Theme"`
	PageSize 	string 	`form:"PageSize"`
}