package models

import "time"

type Todo struct {
	ID         string    `json:"id" gorm:"primary_key;type:char(32)"`
	Content    string    `json:"content" gorm:"type:varchar(255)"`
	Done       bool      `json:"done"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time
	CreatedBy  string
	UpdatedBy  string
	FinishedAt *time.Time `json:"finished_at"`
}
