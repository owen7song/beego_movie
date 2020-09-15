package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(FilmScript))
}

//FilmScript 电影剧本模型
type FilmScript struct {
	//自增id
	ID int64 `orm:"column(id)" json:"id"`
	//频道名称
	Name string `orm:"column(name)" json:"name"`
	//
	Year string `orm:"column(year)" json:"year"`
	//封面
	CoverUrls string `orm:"column(cover_urls)" json:"coverUrls"`
	//创建时间
	CreateTime string `orm:"column(create_time)" json:"createTime"`
}

//TableName 定义表名
func (m *FilmScript) TableName() string {
	return "film_script"
}
