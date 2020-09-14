package repositories

import (
	"beego_movie/models"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

//NewChannelManager 创建频道数据 管理器
func NewChannelManager() *ChannelManager {
	return &ChannelManager{
		TableName: "channel",
	}
}

//ChannelManager 频道管理器
type ChannelManager struct {
	//对应数据库表
	TableName string
}

//Inset 插入一条用户
func (m *ChannelManager) Inset(channel *models.Channel) (ID int64, err error) {
	orm := orm.NewOrm()
	timestamp := time.Now().Unix()
	channel.CreateTime = strconv.FormatInt(timestamp, 10)
	ID, err = orm.Insert(channel)
	return
}

//Select 查询数据
func (m *ChannelManager) Select(page int, pageSize int) (channelList []*models.Channel, count int64, err error) {
	fmt.Println(page, pageSize)
	orm := orm.NewOrm()
	_, err = orm.QueryTable(m.TableName).All(&channelList)
	if err != nil {
		return channelList, 0, err
	}
	count, err = orm.QueryTable(m.TableName).Count()
	if err != nil {
		return channelList, 0, err
	}
	return channelList, count, err
}
