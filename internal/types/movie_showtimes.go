package types

import (
	"time"
)

type MovieShowTime struct {
	TheaterID int       `json:"theater_id"`
	ShowTime  time.Time `json:"show_time"`
}
