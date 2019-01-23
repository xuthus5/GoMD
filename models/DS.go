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
	Count    int    `form:"-"`
	Type     int    `form:"-"` //类型自定义 文章/页面	默认0为文章 1为页面
	Uuid     string `form:"-"`
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

// 评论列表
type Comment struct {
	Id      int    `form:"-"`       //id
	Aid     int    `form:"aid"`     //所属文章id
	Reply   int    `form:"-"`       //对评论的回复
	Content string `form:"content"` //内容
	Date    string `form:"-"`       //时间
	Email   string `form:"email"`   //邮箱
	Name    string `form:"name"`    //名称
	Link    string `form:"link"`    //链接
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
	Date    string //发布时间
	Url     string //定向url
}

//留言 待开发利用
type Message struct {
	Id      int
	Name    string
	Email   string
	Url     string
	Content string
	Date    string
}

// category分类表 慢慢优化完全利用
type Category struct {
	Id          int    //分类方法ID
	Name        string //分类名称
	Key         string //分类标识
	Method      string //分类方法(文章分类article/上传分类upload)
	Description string //分类描述
	Parent      string //所属父分类方法ID
}

type CategoryData struct {
	IsNil 	bool	//check this category is null ?
	Msg 	string
	Info 	[]Category
	List    []Article 	//the article include in this category content
}
// 附件资源记录表
type Attachment struct {
	Id      int    //id
	Name    string //名称
	Type    string `orm:"size(64)"` //分类 独立的文件 属于article的文章
	Path    string //路径
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
type SiteConfigOption struct {
	WebTitle       string `form:"WebTitle"`
	Keywords       string `form:"Keywords"`
	WebUrl         string `form:"WebUrl"`
	Description    string `form:"Description"`
	Status         string `form:"Status"`
	Theme          string `form:"Theme"`
	PageSize       string `form:"PageSize"`
	CopyRight      string `form:"CopyRight"`
	LogoUrl        string `form:"LogoUrl"`
	SecondaryTitle string `form:"SecondaryTitle"`
}

// 用户配置信息
type UserConfigOption struct {
	Author       string `form:"Author"`
	Password     string `form:"Password"`
	UserImageUrl string `form:"UserImageUrl"`
	UserEmail    string `form:"UserEmail"`
	UserQQ       string `form:"UserQQ"`
	UserGithub   string `form:"UserGithub"`
}

// 后台资源管理中心 获取资源分类
type MediaType struct {
	Name string
	Type string
}

/* ---------------------------

功能：前台扩展数据类型

------------------------------ */
