package domain

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	Name         string `json:"name"`
	LastName     string `json:"lastName"`
	MobileNumber string `json:"mobileNumber"`
	UserID       uint
	User         User
	PlanId       uint
	Plan         Plan
}
