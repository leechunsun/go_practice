package repositories

import "chunsheng.isir.com/datamodels"

type MovieRepositories interface {
	GetOneMovie(mid string) *datamodels.Movie
}

type MovieRepositoriesImpl struct {

}

func NewMovieRepositories() MovieRepositories {
	return &MovieRepositoriesImpl{}
}

func (m *MovieRepositoriesImpl) GetOneMovie(mid string) *datamodels.Movie {
	return &datamodels.Movie{"i got a movie " + mid}
}
