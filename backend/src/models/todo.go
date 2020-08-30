package models

import "time"

const (
	StatusNew = iota
	StatusWIP
	StatusDone
	StatusPending
)

type Todo struct {
	ID         string     `json:"id" gorm:"primary_key;type:text"`
	Content    string     `json:"content" gorm:"type:text"`
	Status     int        `json:"status"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"-"`
	FinishedAt *time.Time `json:"finished_at"`
}
