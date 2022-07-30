package domain

import "gorm.io/gorm"

type Withdraw struct {
	gorm.Model
	Card   Card
	Amount int
	CardId uint
}
