package service

import (
	"fmt"

	"github.com/GareArc/MovieMate/internal/db"
	"github.com/GareArc/MovieMate/internal/types"
	"github.com/GareArc/MovieMate/internal/types/model"
	"github.com/rs/zerolog/log"
)

type MovieService struct{}

func (ms *MovieService) GetMovieInfoById(id string) (*model.Movie, error) {
	db := db.MainDB
	var movie = &model.Movie{}
	db.Model(&model.Movie{}).Where("id = ?", id).First(&movie)
	if movie.ID == "" {
		log.Error().Msgf("movie with id=%s not found", id)
		return nil, fmt.Errorf("movie with id=%s not found", id)
	}

	return movie, nil
}

func (ms *MovieService) GetShowTimeListById(id string) []types.MovieShowTime {
	db := db.MainDB

	var schedules []model.MovieSchedule
	db.Model(&model.MovieSchedule{}).Where("movie_id = ?", id).Find(&schedules)
	if (len(schedules)) == 0 {
		log.Info().Msgf("movie with id %s does not have any showtime", id)
	}

	showtime_list := make([]types.MovieShowTime, len(schedules))
	for i, schedule := range schedules {
		showtime_list[i] = types.MovieShowTime{
			TheaterID: schedule.TheaterID,
			ShowTime:  schedule.ShowTime,
		}
	}

	return showtime_list
}
