package repositories

import (
	"beego_movie/models"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

//NewFilmScriptManager 创建数据管理器
func NewFilmScriptManager() *FilmScriptManager {
	return &FilmScriptManager{
		TableName: "film_script",
	}
}

//FilmScriptManager 电影剧本数据管理器
type FilmScriptManager struct {
	//对应数据库表
	TableName string
}

//Inset 插入一条数据
func (m *FilmScriptManager) Inset(filmScript *models.FilmScript) (ID int64, err error) {
	orm := orm.NewOrm()
	timestamp := time.Now().Unix()
	filmScript.CreateTime = strconv.FormatInt(timestamp, 10)
	ID, err = orm.Insert(filmScript)
	return
}

//Update 更新数据
func (m *FilmScriptManager) Update(filmScript *models.FilmScript) bool {
	orm := orm.NewOrm()
	if orm.Read(filmScript) == nil {
		if _, err := orm.Update(filmScript); err == nil {
			return true
		}
	}
	return false
}
