package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(Channel))
}

//Channel 数据模型
type Channel struct {
	//自增id
	ID int64 `orm:"column(id)" json:"id"`
	//频道名称
	Name string `orm:"column(channel_name)" json:"name"`
	//创建时间
	CreateTime string `orm:"column(create_time)" json:"createTime"`
}

//TableName 定义表名
func (m *Channel) TableName() string {
	return "channel"
}
