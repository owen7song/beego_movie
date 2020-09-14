package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//InitDB 初始化数据库
func InitDB() {
	db := beego.AppConfig.String("db")
	err := orm.RegisterDataBase("default", "mysql", db)
	if err != nil {
		panic("mysql connect failed, detail is " + err.Error() + "")
	}

	fmt.Println("mysql connect success")
}
