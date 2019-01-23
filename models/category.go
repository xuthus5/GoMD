package models

import "strconv"

/*********

the model about category

*********/

//return the category information
func CategoryInformation(fields string) *CategoryData{
	info := []Category{}
	data := CategoryData{}
	dbx.Select(&info, "select * from category where key=?", fields)
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