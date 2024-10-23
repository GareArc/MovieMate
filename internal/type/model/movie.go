package model

import (
	"time"
)

type ClassAllowed string

const (
	ALLAGES  ClassAllowed = "all_ages"
	TEENAGER ClassAllowed = "teenager"
	ADULT    ClassAllowed = "adult"
)

type Movie struct {
	BaseModel
	Name              string       `gorm:"column:name;length:255;not null"`
	Description       string       `gorm:"column:description;length:255"`
	Poster            string       `gorm:"column:poster;length:255"`
	ReleaseDate       string       `gorm:"column:release_date;length:255"`
	DurationInMinutes int          `gorm:"column:duration;default:0"`
	Language          string       `gorm:"column:language;length:255"`
	Subtitle          string       `gorm:"column:subtitle;length:255"`
	IMax              bool         `gorm:"column:imax;default:false"`
	Class             ClassAllowed `gorm:"column:class;length:255;default:'all_ages'"`
}

type MoiveSchedule struct {
	BaseModel
	MovieID   int       `gorm:"column:movie_id;not null"`
	TheaterID int       `gorm:"column:theater_id;not null"`
	ShowTime  time.Time `gorm:"column:show_time;not null"`
}
