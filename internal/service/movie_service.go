package service

import (
	"fmt"

	"github.com/GareArc/MovieMate/internal/db"
	"github.com/GareArc/MovieMate/internal/type/model"
	"github.com/rs/zerolog/log"
)

type MovieService struct{}

func (ms *MovieService) GetMovieInfoById(id string) (*model.Movie, error) {
	db := db.MainDB
	var movie = &model.Movie{}
	db.Model(&model.Movie{}).Where("id = ?", id).First(movie)
	if movie.ID == "" {
		log.Error().Msgf("movie with id=%s not found", id)
		return nil, fmt.Errorf("movie with id=%s not found", id)
	}

	return movie, nil
}
