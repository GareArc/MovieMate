package model

type Seat struct {
	BaseModel
	Row       int    `gorm:"column:row;not null"`
	Col       string `gorm:"column:col;length:255;not null"`
	Avaliable bool   `gorm:"column:avaliable;default:true"`
	TheaterID int    `gorm:"column:theater_id;not null"`
}
