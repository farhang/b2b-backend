package domain

import "gorm.io/gorm"

type Deposit struct {
	gorm.Model
	UserId uint
	User   User
}
