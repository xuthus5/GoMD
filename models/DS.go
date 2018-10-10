package models

/* ---------------------------

功能：database structure
数据库，数据类型，字段统一存放位置

------------------------------ */

/* ---------------------------

Type：database structure

------------------------------ */

// 全局配置表  Option=配置选项  Value=配置内容
type Config struct {
	Id     int
	Option string `orm:"size(16)"`
	Value  string `orm:"size(32)"`
}

//文章表 存储完整的文章信息 并映射文章提交的数据字段
type Article struct {
	Id       int    `form:"-"`
	Cid      int    `form:"cid"`
	Type     int    `form:"-"`
	Title    string `orm:"size(64)" form:"title"`
	Author   string `form:"author"`
	Image    string `form:"image"`
	Tags     string `orm:"size(64)" form:"tags"`
	Renew    string `orm:"size(20)" form:"-"`
	Content  string `orm:"type(text)" form:"content"`
	Summary  string `form:"summary"`
	Status   string `form:"status"`
	Password string `form:"password"`
	Link     string `form:"-"`
}

// 链接表
type Link struct {
	Id          int    //自增唯一ID
	Url         string //链接URL
	Name        string //链接标题
	Image       string //链接图片
	Target      string //链接打开方式
	Description string //链接描述
	Visible     string //是否可见（Y/N）
}

//公告表
type Notice struct {
	Id      int    //自增唯一ID
	Content string //内容
	Data    string //发布时间
	Url     string //定向url
}

//留言 待开发利用
type Message struct {
	Id      int
	Name    string
	Email   string
	Url     string
	Content string
	Data    string
}

// taxonomy分类表 慢慢优化完全利用
type Taxonomy struct {
	Id          int    //分类方法ID
	Name        string //分类名称
	Key         string //分类标识
	Method      string //分类方法(文章分类article/上传分类upload)
	Description string //分类描述
	Parent      string //所属父分类方法ID
}

// 资源记录表
type Source struct {
	Id      int                     //id
	Name    string                  //名称
	Type    string `orm:"size(64)"` //分类
	Path    string                  //路径
	Created string `orm:"size(10)"` //创建时间
}

/* ---------------------------

功能：后台扩展数据类型

------------------------------ */

// 在后台显示的文章操作界面 数据结构
type DisplayArticle struct {
	Id     int
	Title  string
	Author string
	Name   string
	Tags   string
	Renew  string
}


// 网站后台提交的表单字段 映射到此结构体 需要持续添加
type ConfigOption struct {
	WebTitle    string `form:"WebTitle"`
	Keywords    string `form:"Keywords"`
	WebUrl      string `form:"WebUrl"`
	Description string `form:"Description"`
	Status      string `form:"Status"`
	Theme       string `form:"Theme"`
	PageSize    string `form:"PageSize"`
}

// 后台资源管理中心 获取资源分类
type MediaType struct {
	Name string
	Type string
}

/* ---------------------------

功能：前台扩展数据类型

------------------------------ */
