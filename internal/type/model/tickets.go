package model

type Ticket struct {
	BaseModel
	SeatID          int `gorm:"column:seat_id;not null"`
	MovieScheduleID int `gorm:"column:movie_schedule_id;not null"`
	Price           int `gorm:"column:price;not null"`
}

type TicketPurchase struct {
	BaseModel
	TicketID int    `gorm:"column:ticket_id;not null"`
	UserID   int    `gorm:"column:user_id;not null"`
	Status   string `gorm:"column:status;length:255;default:'pending'"`
}
