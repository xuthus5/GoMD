package models

/* ---------------------------------

功能： 文章处理模块 包含文章表，分类表的处理
	  网站架构的处理中心

------------------------------------*/

/*
	AllArticleList() 获取数据表所有文章
	GetOneArticle()	 查看某一篇文章 需要提供ID
	LimitArticleDisplay() 限制文章数量  用于文章分页
	Search() 	搜索文章功能
	DeleteArticle	删除
	UpdateArticle	修改
	AddArticle	添加

*/
// 返回数据表所有文章
func AllArticleList() (*[]Article, error) {
	list := []Article{}
	err := dbx.Select(&list, "select * from article order by id desc")
	return &list, err
}

// 查看某一篇文章 提供ID 查询文章
func GetOneArticle(id, method string) *[]Article {
	data := []Article{}
	if method == "uuid" {
		dbx.Select(&data, "select * from article where uuid=?", id)
	} else {
		dbx.Select(&data, "select * from article where id=?", id)
	}
	return &data
}

// 查看某一篇文章 提供ID 查询文章
func GetPreOrNextArticle(id int, method string) *map[string]string {
	data := []Article{}
	res := make(map[string]string)
	if method == "pre" {
		dbx.Select(&data, "select * from article where id = (select id from article where id < ? order by id desc limit 1);", id)
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
		dbx.Select(&data, "select * from article where id = (select id from article where id > ? order by id asc limit 1);", id)
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
	err := dbx.Select(&list, "select * from article order by id desc limit ?,?", page, limit)
	if err != nil {
		panic(err.Error())
	}
	count, _ := dbc.QueryTable("article").Count()
	return &list, count
}

// 文章搜索功能 返回搜索列表
func Search(keywords string) (*[]Article, error) {
	list := []Article{}
	err := dbx.Select(&list, "select * from article where `title`||`content` like '%"+keywords+"%'")
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
	dbc.Read(article)
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
		panic(err.Error())
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
	err := dbx.Select(&data, "select * from comment where aid=?", id)
	if err != nil {
		panic(err.Error())
	}
	return &data
}

// 得到一篇文章下的评论数量 传入文章id
func GetCommentNumFromArticle(id int) (int64, error) {
	qs := dbc.QueryTable("comment")
	return qs.Filter("aid", id).Count()
}

// 根据传递过来的id返回uuid
func GetUuidById(id int) string {
	qs := dbc.QueryTable("article")
	var article *Article
	qs.Filter("id", id).One(article, "uuid")
	return article.Uuid
}

// ---------------分类管理----------------------
/*
	AddCategory	添加一个分类
	CategoryList	查询所有的分类
*/

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

// 查询分类列表  用于后台展示所有分类
func CategoryList() *[]Category {
	list := []Category{}
	err := dbx.Select(&list, "select * from category")
	if err != nil {
		panic(err.Error())
	}
	return &list
}

// 获得一个分类信息 主要用于更新分类信息
func GetOneCategoryInfo(key, method string) *[]Category {
	list := []Category{}
	if method == "id" {
		dbx.Select(&list, "select * from category where id=?", key)
	} else {
		dbx.Select(&list, "select * from category where key=?", key)
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
	dbc.Read(data)
	_, err := dbc.Delete(data)
	return data.Path, err
}

// 返回一个附件信息
func FileInfo(id int64) *[]Attachment {
	data := []Attachment{}
	err := dbx.Select(&data, "select * from attachment where id=?", id)
	if err != nil {
		panic(err.Error())
	}
	return &data
}
