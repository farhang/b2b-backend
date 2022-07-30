package domain

import "gorm.io/gorm"

type Asset struct {
	gorm.Model
	Tether      int `json:"tether_balance"`
	InPlan      int `json:"in_plan"`
	CanWithdraw int `json:"can_withdraw"`
	UserID      uint
	User        User
}
