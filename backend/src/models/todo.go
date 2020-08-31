package models

import "time"

type TodoStatus int

const (
	StatusNew TodoStatus = iota
	StatusWIP
	StatusDone
	StatusPending
	StatusUnknown
)

func (ts TodoStatus) String() string {
	switch ts {
	case StatusNew:
		return "new"
	case StatusWIP:
		return "wip"
	case StatusDone:
		return "done"
	case StatusPending:
		return "pending"
	default:
		return "unknown"
	}
}

func NewTodoStatus(status int) TodoStatus {
	if status == 0 {
		return StatusNew
	}
	if status == 1 {
		return StatusWIP
	}
	if status == 2 {
		return StatusDone
	}
	if status == 3 {
		return StatusPending
	}
	return StatusUnknown
}

type Todo struct {
	ID         string     `json:"id" gorm:"primary_key;type:text"`
	Content    string     `json:"content" gorm:"type:text"`
	Status     TodoStatus `json:"status"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"-"`
	FinishedAt *time.Time `json:"finished_at"`
}
