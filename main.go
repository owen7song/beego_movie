package main

import (
	"beego_movie/models"
	_ "beego_movie/routers"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	models.InitDB()
}

func main() {
	beego.Run()
}
