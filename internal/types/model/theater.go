package model

type Theater struct {
	BaseModel
	Name      string `gorm:"column:name;length:255;not null"`
	Avaliable bool   `gorm:"column:avaliable;default:true"`
	TotalSeat int    `gorm:"column:total_seat;default:0"`
}
