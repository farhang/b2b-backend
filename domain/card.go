package domain

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	Sheba  string
	UserId uint
	User   User
}
