package domain

import "gorm.io/gorm"

type Plan struct {
	gorm.Model
	Title         string `json:"title"`
	Description   string `json:"description"`
	ProfitPercent int    `json:"profit_percent"`
	Duration      int    `json:"duration_in_month"`
}
