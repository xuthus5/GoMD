package models

import "log"

/* ---------------------------------

功能： 公共操作函数

------------------------------------*/

/* ---------------------------------
配置模块
------------------------------------*/
// 获取所有配置信息 用于首页的网站信息显示
func ConfigList() map[string]string {
	list := []Config{}
	err := dbx.Select(&list, "select * from config")
	if err != nil {
		log.Println(err.Error())
	}
	var data = make(map[string]string)
	for _, v := range list {
		data[v.Option] = v.Value
	}
	return data
}

// 网站全局配置 用于更新一条网站配置信息
func SiteConfig(data *Config) error {
	raw := new(Config)
	raw = &Config{Option: data.Option, Value: data.Value}
	if _, _, err := dbc.ReadOrCreate(data, "Option"); err != nil {
		return err
	} else {
		_, err := dbc.Raw("UPDATE config SET value = ? WHERE option = ?", raw.Value, raw.Option).Exec()
		return err
	}
}

/* ---------------------------------
用户管理模块
------------------------------------*/

// 用户登陆校验
func Login(username, password *Config) error {
	user := dbc.Read(username, "Option", "Value")
	pass := dbc.Read(password, "Option", "Value")
	if user != nil {
		return user
	}
	if pass != nil {
		return pass
	}
	return nil
}

/* ---------------------------------
其他模块
------------------------------------*/

// 返回表中的数据条数 用于数据统计
func TableNumber(table string) (count int64, err error) {
	count, err = dbc.QueryTable(table).Count()
	return
}

// 用于获取网站配置信息
func GetOneConfig(info string) string {
	data := []Config{}
	err := dbx.Select(&data, "select * from config where Option=?", info)
	if err != nil {
		log.Println(err.Error())
	}
	return data[0].Value
}

//公告管理模块
func AddNotice(data *Notice) error {
	_, err := dbc.Insert(data)
	return err
}

//返回一条最近公告
func GetOneNotice() string {
	data := []Notice{}
	err := dbx.Select(&data, "select * from notice order by date desc limit 0,1")
	if err != nil {
		log.Println(err.Error())
	}
	return data[0].Content
}
