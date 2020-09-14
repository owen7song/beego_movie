package routers

import (
	"beego_movie/controllers/admin"
	"beego_movie/middleware"

	"github.com/astaxie/beego"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, middleware.InitCors)
	beego.Router("/admin/channel/create", admin.NewChnnelController(), "POST:Create")
	beego.Router("/admin/channel/select", admin.NewChnnelController(), "GET:SelectChnnel")
}
