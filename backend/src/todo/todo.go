package todo

import "time"

type TODO struct {
	ID         string    `json:"id"`
	Content    string    `json:"content"`
	Done       bool      `json:"done"`
	CreatedAt  time.Time `json:"created_at"`
	FinishedAt time.Time `json:"finished_at"`
}
