package admin

import (
	"beego_movie/common"
	"beego_movie/models"
	"beego_movie/service"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

//NewChnnelController 创建频道控制器
func NewChnnelController() *ChnnelController {
	ChannelService := service.NewChannelService()
	return &ChnnelController{
		ChannelService: ChannelService,
	}
}

//ChnnelController 频道管理器
type ChnnelController struct {
	beego.Controller
	ChannelService *service.ChannelService
}

//Create 创建频道 *POST
func (c *ChnnelController) Create() {
	var ob models.Channel
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &ob)
	ok := c.ChannelService.CreateChannel(&ob)
	if ok {
		c.Data["json"] = common.ResponseSuccess(nil)
	} else {
		c.Data["json"] = common.ResponseFailed(nil)
	}
	c.ServeJSON()
}

// SelectChnnel 查询频道 *GET
func (c *ChnnelController) SelectChnnel() {
	channelList, count, err := c.ChannelService.QeuryChannel(1, 10)
	if err != nil {
		c.Data["json"] = common.ResponseFailed(nil)
	} else {
		c.Data["json"] = common.ResponseSuccess(channelList)
		fmt.Println(count)
	}
	c.ServeJSON()
}
