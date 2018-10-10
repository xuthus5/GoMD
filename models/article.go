package models

/* ---------------------------------

功能： 文章处理模块 包含文章表，分类表的处理

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
func AllArticleList() (*[]Article,error){
	list := []Article{}
	err := dbx.Select(&list, "select * from article order by id desc")
	return &list,err
}
// 查看某一篇文章 提供ID 查询文章
func GetOneArticle(id string) (*[]Article,error){
	data := []Article{}
	err := dbx.Select(&data, "select * from article where id=?",id)
	return &data,err
}
// 限制文章数量  用于文章分页
func LimitArticleDisplay(page,limit int64) (*[]Article,int64){
	list := []Article{}
	page = limit*(page-1)
	err := dbx.Select(&list, "select * from article order by id desc limit ?,?",page,limit)
	if err != nil {
		panic(err.Error())
	}
	count,_ := dbc.QueryTable("article").Count()
	return &list,count
}
// 文章搜索功能 返回搜索列表
func Search(keywords string) (*[]Article,error){
	list := []Article{}
	err := dbx.Select(&list, "select * from article where `title`||`content` like '%"+keywords+"%'")
	return &list,err
}
// 文章更新
func UpdateArticle(data *Article) error {
	_, err := dbc.Update(data)
	if err != nil {
		return err
	}else{
		return nil
	}
}
// 添加文章
func AddArticle(data *Article) error{
	_, err := dbc.Insert(data)
	return err
}
// 文章删除
func DeleteArticle(id int) error {
	data := &Article{Id:id}
	_,err := dbc.Delete(data)
	return err
}
// 查询一篇文章的分类
func SearchArticleCategory(id int) *[]Taxonomy{
	data := []Taxonomy{}
	err := dbx.Select(&data, "select * from taxonomy where id=?",id)
	if err != nil {
		panic(err.Error())
	}
	return &data
}


// ---------------分类管理----------------------
/*
	AddCategory	添加一个分类
	CategoryList	查询所有的分类
 */

// 添加分类
func AddCategory(data *Taxonomy) error{
	_, err := dbc.Insert(data)
	return err
}
// 分类删除
func DeleteCategory(id int) error {
	data := &Taxonomy{Id:id}
	_,err := dbc.Delete(data)
	return err
}
// 分类修改
func UpdateCategory(data *Taxonomy) error{
	_, err := dbc.Update(data)
	if err != nil {
		return err
	}else{
		return nil
	}
}
// 查询分类列表  用于后台展示所有分类
func CategoryList() *[]Taxonomy{
	list := []Taxonomy{}
	err := dbx.Select(&list, "select * from taxonomy")
	if err != nil {
		panic(err.Error())
	}
	return &list
}
// 获得一个分类信息 主要用于更新分类信息
func GetOneCategoryInfo(id string) *[]Taxonomy{
	list := []Taxonomy{}
	err := dbx.Select(&list, "select * from taxonomy where id=?",id)
	if err != nil {
		panic(err.Error())
	}
	return &list
}