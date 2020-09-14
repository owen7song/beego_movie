package service

import (
	"beego_movie/models"
	"beego_movie/repositories"
	"fmt"
)

//NewChannelService 创建频道 服务
func NewChannelService() *ChannelService {
	channelManager := repositories.NewChannelManager()
	return &ChannelService{
		ChannerReposioty: channelManager,
	}
}

//ChannelService 频道服务
type ChannelService struct {
	ChannerReposioty *repositories.ChannelManager
}

//CreateChannel 创建频道
func (s *ChannelService) CreateChannel(channel *models.Channel) bool {
	_, err := s.ChannerReposioty.Inset(channel)
	if err != nil {
		fmt.Println("频道创建失败")
		return false
	}
	return true
}

//QeuryChannel 查询频道
func (s *ChannelService) QeuryChannel(page int, pageSize int) (channelList []*models.Channel, count int64, err error) {
	return s.ChannerReposioty.Select(page, pageSize)
}
