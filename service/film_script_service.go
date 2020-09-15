package service

import (
	"beego_movie/models"
	"beego_movie/repositories"
	"fmt"
)

//NewFilmScriptService 创建服务实列
func NewFilmScriptService() *FilmScriptService {
	filmScriptManager := repositories.NewFilmScriptManager()
	return &FilmScriptService{
		FilmScriptRepository: filmScriptManager,
	}
}

//FilmScriptService 电影剧本服务
type FilmScriptService struct {
	FilmScriptRepository *repositories.FilmScriptManager
}

//CreateFilmScript 创建电影剧本
func (s *FilmScriptService) CreateFilmScript(filmScript *models.FilmScript) bool {
	_, err := s.FilmScriptRepository.Inset(filmScript)
	if err != nil {
		fmt.Println("创建电影剧本失败")
		return false
	}
	return true
}
