package domain

import "gorm.io/gorm"

type Credit struct {
	gorm.Model
	Amount int
	User   User
	UserId uint
}
