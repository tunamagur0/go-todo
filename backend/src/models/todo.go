package models

import "time"

type Todo struct {
	ID         string    `json:"id" gorm:"primary_key;type:text"`
	Content    string    `json:"content" gorm:"type:text"`
	Done       bool      `json:"done"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time
	FinishedAt *time.Time `json:"finished_at"`
}
