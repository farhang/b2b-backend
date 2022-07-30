package domain

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	Profit       int  `json:"profit"`
	IsDescending bool `json:"is_descending"`
	IsAscending  bool `json:"is_ascending"`
	UserId       uint
	User         User
}
