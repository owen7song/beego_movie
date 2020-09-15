package admin

import (
	"beego_movie/common"
	"beego_movie/models"
	"beego_movie/service"
	"encoding/json"

	"github.com/astaxie/beego"
)

//NewFilmScriptController 剧本
func NewFilmScriptController() *FilmScriptController {
	filmScriptService := service.NewFilmScriptService()
	return &FilmScriptController{
		FilmScriptService: filmScriptService,
	}
}

//FilmScriptController 电影剧本控制器
type FilmScriptController struct {
	beego.Controller
	FilmScriptService *service.FilmScriptService
}

//CreateFilm 创建电影
func (c *FilmScriptController) CreateFilm() {
	var ob models.FilmScript
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &ob)
	ok := c.FilmScriptService.CreateFilmScript(&ob)
	if ok {
		c.Data["json"] = common.ResponseSuccess(nil)
	} else {
		c.Data["json"] = common.ResponseFailed(nil)
	}
	c.ServeJSON()
}
