package model

type Account struct {
	Model
	Email        string `gorm:"column:email;unique;not null"`
	PasswordHash string `gorm:"column:password;not null"`
	Salt         string `gorm:"column:salt;not null"`
}
