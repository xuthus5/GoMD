package models

import (
	"GoMD/tools"
	"log"
	"strconv"
)

/* ---------------------------------

功能： 网站架构的处理中心

------------------------------------*/

// ---------------文章管理----------------------

// 返回数据表所有文章
func AllArticleList() (*[]Article, error) {
	list := []Article{}
	err := dbx.Select(&list, "select * from article where type=0 order by id desc")
	return &list, err
}

// 查看某一篇文章 提供ID 查询文章
func GetOneArticle(id, method string) *[]Article {
	data := []Article{}
	if method == "uuid" {
		_ = dbx.Select(&data, "select * from article where uuid=?", id)
	} else {
		_ = dbx.Select(&data, "select * from article where id=?", id)
	}
	return &data
}

// 提供ID 查看文章的属性
func GetPropertyByID(id int,Property string) string {
	data := []Article{}
	_ = dbx.Select(&data, "select * from article where id=?", id)
	if Property == "title" {
		return data[0].Title
	}
	return ""
}

// 查看某一篇文章的上下文 提供ID
func GetPreOrNextArticle(id int, method string) *map[string]string {
	data := []Article{}
	res := make(map[string]string)
	if method == "pre" {
		_ = dbx.Select(&data, "select * from article where id = (select id from article where id < ? order by id desc limit 1);", id)
		if len(data) == 0 {
			res["isNull"] = "true"
			res["title"] = ""
			res["uuid"] = ""
		} else {
			res["isNull"] = "false"
			res["title"] = data[0].Title
			res["uuid"] = data[0].Uuid
		}
	} else {
		_ = dbx.Select(&data, "select * from article where id = (select id from article where id > ? order by id asc limit 1);", id)
		if len(data) == 0 {
			res["isNull"] = "true"
			res["title"] = ""
			res["uuid"] = ""
		} else {
			res["isNull"] = "false"
			res["title"] = data[0].Title
			res["uuid"] = data[0].Uuid
		}
	}
	return &res
}

// 限制文章数量  用于文章分页
func LimitArticleDisplay(page, limit int64) (*[]Article, int64) {
	list := []Article{}
	page = limit * (page - 1)
	err := dbx.Select(&list, "select * from article where type=0 order by id desc limit ?,?", page, limit)
	if err != nil {
		log.Println(err.Error())
	}
	count, _ := dbc.QueryTable("article").Count()
	return &list, count
}

// 返回指定条数最新文章
func GetLimitNewArticle(num int64) *[]Article {
	list := []Article{}
	_ = dbx.Select(&list, "select * from article where type=0 order by renew desc limit 0,?", num)
	return &list
}

// 文章搜索功能 返回搜索列表
func Search(keywords string) (*[]Article, error) {
	list := []Article{}
	err := dbx.Select(&list, "select * from article where type=0 and `title`||`content`||`tags` like '%"+keywords+"%'")
	return &list, err
}

// 文章更新
func UpdateArticle(data *Article) error {
	_, err := dbc.Update(data)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// 添加文章
func AddArticle(data *Article) (string, error) {
	id, err := dbc.Insert(data)
	article := &Article{Id: int(id)}
	_ = dbc.Read(article)
	return article.Uuid, err
}

// 文章删除
func DeleteArticle(id int) error {
	data := &Article{Id: id}
	_, err := dbc.Delete(data)
	return err
}

// 查询一篇文章的分类
func SearchArticleCategory(id int) *[]Category {
	data := []Category{}
	err := dbx.Select(&data, "select * from category where id=?", id)
	if err != nil {
		log.Println(err.Error())
	}
	return &data
}

// ---------------评论管理----------------------

// 添加一条评论
func AddComment(data *Comment) error {
	_, err := dbc.Insert(data)
	return err
}

// 得到一篇文章下的所有评论
func GetArticleComments(id int) *[]Comment {
	data := []Comment{}
	err := dbx.Select(&data, "select * from comment where status=1 and aid=?", id)
	if err != nil {
		log.Println(err.Error())
	}
	return &data
}

// 得到一篇文章下的评论数量 传入文章id
func GetCommentNumFromArticle(id int) (int64, error) {
	return dbc.QueryTable("comment").Filter("aid", id).Filter("status",1).Count()
}

// 根据传递过来的id返回uuid
func GetUuidById(id int64) string {
	var article Article
	_ = dbc.QueryTable("article").Filter("id", id).One(&article, "uuid")
	return article.Uuid
}

// 返回指定条数最新评论
func GetLimitNewComment(num int64) *[]Comment {
	list := []Comment{}
	_ = dbx.Select(&list, "select * from comment order by date desc limit 0,?", num)
	return &list
}

// 删除文章下的所有评论
func DeleteCommentFromArticle(id int) error {
	data := &Comment{Aid: id}
	_, err := dbc.Delete(data, "aid")
	return err
}

// ---------------分类管理----------------------

// 添加分类
func AddCategory(data *Category) error {
	_, err := dbc.Insert(data)
	return err
}

// 分类删除
func DeleteCategory(id int) error {
	//删除分类的同时，需要将旗下的文章删除
	data := &Category{Id: id}
	_, err := dbc.Delete(data)
	if err != nil {
		return err
	}
	article := &Article{Cid: id}
	_, err = dbc.Delete(article, "Cid")
	if err != nil {
		return err
	}
	return nil
}

// 分类修改
func UpdateCategory(data *Category) error {
	_, err := dbc.Update(data)
	if err != nil {
		return err
	} else {
		return nil
	}
}


// 返回分类的信息
func CategoryInformation(fields string) *CategoryData{
	info := []Category{}
	data := CategoryData{}
	_ = dbx.Select(&info, "select * from category where key=?", fields)
	if len(info)==0{
		data.IsNil = true
		data.Msg = "分类不存在!"
	}else{
		data.IsNil = false
		data.Info = info
		list,_ := GetCategoryArticle(strconv.Itoa(info[0].Id))
		data.List = *list
	}
	return &data
}

// 查询分类列表  用于后台展示所有分类
func CategoryList() *[]Category {
	list := []Category{}
	err := dbx.Select(&list, "select * from category")
	if err != nil {
		log.Println(err.Error())
	}
	return &list
}

// 获得一个分类信息 主要用于更新分类信息
func GetOneCategoryInfo(key, method string) *[]Category {
	list := []Category{}
	if method == "id" {
		_ = dbx.Select(&list, "select * from category where id=?", key)
	} else {
		_ = dbx.Select(&list, "select * from category where key=?", key)
	}
	return &list
}

// 返回分类下的所有文章信息
func GetCategoryArticle(cid string) (*[]Article, error) {
	list := []Article{}
	err := dbx.Select(&list, "select * from article where cid= ? order by id desc", cid)
	return &list, err
}

//返回一个分类下的文章数量
func GetNumFromACategory(cid int) (int64, error) {
	qs := dbc.QueryTable("article")
	return qs.Filter("cid", cid).Count()
}

// ---------------附件管理----------------------

//文件上传 入数据库
func FileSave(info *Attachment) (int64, error) {
	return dbc.Insert(info)
}

// 文件删除
func FileDelete(id int) (string, error) {
	data := &Attachment{Id: id}
	_ = dbc.Read(data)
	_, err := dbc.Delete(data)
	return data.Path, err
}

// 返回一个附件信息
func FileInfo(id int64) *[]Attachment {
	data := []Attachment{}
	err := dbx.Select(&data, "select * from attachment where id=?", id)
	if err != nil {
		log.Println(err.Error())
	}
	return &data
}

//--------------------链接管理-------------------

//添加链接
func AddLink(info *Link) error {
	_, err := dbc.Insert(info)
	return err
}

//删除链接
func DeleteLink(id int) error {
	data := &Link{Id: id}
	_ = dbc.Read(data)
	_, err := dbc.Delete(data)
	return err
}

//获取所有链接list
func GetAllLink() *[]Link {
	list := []Link{}
	_ = dbx.Select(&list, "select * from link where type=0 order by id desc")
	return &list
}

// 获取一个link 信息
func GetOneLinkInfo(id string) *[]Link {
	data := []Link{}
	err := dbx.Select(&data, "select * from link where id=?", id)
	if err != nil {
		log.Println(err.Error())
	}
	return &data
}

// 修改链接信息
func UpdateLink(data *Link) error {
	_, err := dbc.Update(data)
	if err != nil {
		return err
	} else {
		return nil
	}
}

//--------------------菜单栏管理-------------------
//添加链接
func AddMenuNode(info *Link) error {
	_, err := dbc.Insert(info)
	return err
}

//删除链接
func DeleteMenuNode(id int) error {
	data := &Link{Id: id}
	_ = dbc.Read(data)
	_, err := dbc.Delete(data)
	return err
}

//获取所有链接list
func GetAllMenu() *[]Link {
	list := []Link{}
	_ = dbx.Select(&list, "select * from link where type=1 order by id desc")
	return &list
}

// 获取一个link 信息
func GetOneMenuNodeInfo(id string) *[]Link {
	data := []Link{}
	err := dbx.Select(&data, "select * from link where id=?", id)
	if err != nil {
		log.Println(err.Error())
	}
	return &data
}

// 修改链接信息
func UpdateMenuNode(data *Link) error {
	_, err := dbc.Update(data)
	if err != nil {
		return err
	} else {
		return nil
	}
}

//--------------------评论管理-------------------

// 待审核评论列表
func ReviewComment() *[]DisplayComment {
	list := []DisplayComment{}
	err := dbx.Select(&list, "select comment.content,comment.date,comment.name,comment.id,article.title from comment,article where article.id=comment.aid and comment.status=0")
	if err != nil {
		log.Println(err.Error())
	}
	for i, v := range list {
		list[i].Date = tools.UnixTimeToString(v.Date)
	}
	return &list
}

// 已审核评论列表
func AdoptComment() *[]DisplayComment {
	list := []DisplayComment{}
	err := dbx.Select(&list, "select comment.content,comment.date,comment.name,comment.id,article.title from comment,article where article.id=comment.aid and comment.status=1")
	if err != nil {
		log.Println(err.Error())
	}
	for i, v := range list {
		list[i].Date = tools.UnixTimeToString(v.Date)
	}
	return &list
}

// 评论删除
func CommentDelete(id int) error {
	data := &Comment{Id: id}
	_ = dbc.Read(data)
	_, err := dbc.Delete(data)
	return err
}

//修改评论状态
func CommentUpdate(data *Comment) error{
	_, err := dbc.Update(data,"Status")
	if err != nil {
		return err
	} else {
		return nil
	}
}