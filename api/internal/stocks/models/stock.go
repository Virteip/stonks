package models

import (
	"time"
)

// Stock represents stock information for our domain
type Stock struct {
	ID         string    `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Ticker     string    `json:"ticker" gorm:"size:10;not null"`
	Company    string    `json:"company" gorm:"size:255;not null"`
	Brokerage  string    `json:"brokerage" gorm:"size:255;not null"`
	Action     string    `json:"action" gorm:"size:50;not null"`
	RatingFrom string    `json:"rating_from" gorm:"size:50"`
	RatingTo   string    `json:"rating_to" gorm:"size:50"`
	TargetFrom float64   `json:"target_from"`
	TargetTo   float64   `json:"target_to"`
	Time       time.Time `json:"time" gorm:"type:timestamp;not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:timestamp;autoUpdateTime"`
}

// PaginatedStocks represents paginated stock data
type PaginatedStocks struct {
	Stocks     []Stock `json:"stocks"`
	TotalCount int64   `json:"total_count"`
	PageSize   int     `json:"page_size"`
	Page       int     `json:"page"`
	TotalPages int     `json:"total_pages"`
}

type PaginationParams struct {
	Page     int
	PageSize int
}
