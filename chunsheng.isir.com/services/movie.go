package services

import "chunsheng.isir.com/repositories"

type MovieService interface {
	GetMovieName(mid string) string
}

type MovieServiceImpl struct {
	Res repositories.MovieRepositories
}


func NewMovieService(res repositories.MovieRepositories) MovieService {
	return &MovieServiceImpl{res}
}

func (m *MovieServiceImpl) GetMovieName(mid string) string {
	return m.Res.GetOneMovie(mid).Name
}

